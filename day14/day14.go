package day14

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type Chemical struct {
	Name           string
	Inputs         map[string]int
	OutputQuantity int
}

func NewChemical(formula string) Chemical {
	parsed := strings.Split(formula, " => ")
	inputsMap := map[string]int{}
	for _, input := range strings.Split(parsed[0], ", ") {
		parsedInput := strings.Split(input, " ")
		inputQty, err := strconv.Atoi(parsedInput[0])
		if err != nil {
			log.Fatalf("error parsing quantity: %v", err)
		}
		inputName := parsedInput[1]
		inputsMap[inputName] = inputQty
	}

	parsedOutput := strings.Split(parsed[1], " ")
	outputQty, err := strconv.Atoi(parsedOutput[0])
	if err != nil {
		log.Fatalf("error parsing quantity: %v", err)
	}
	outputName := parsedOutput[1]

	return Chemical{Inputs: inputsMap, OutputQuantity: outputQty, Name: outputName}
}

// returns the amount of input material required and the number of excess it produces
func (chemical Chemical) Make(qty int) (map[string]int, int) {
	inputs := map[string]int{}
	reactionsNeeded := int(math.Ceil(float64(qty) / float64(chemical.OutputQuantity)))
	for inputChemical, inputQty := range chemical.Inputs {
		inputs[inputChemical] = reactionsNeeded * inputQty
	}

	return inputs, (chemical.OutputQuantity * reactionsNeeded) - qty
}

type Chemicals map[string]Chemical

func NewChemicals(specs string) Chemicals {
	chemicals := Chemicals{}
	for _, spec := range strings.Split(specs, "\n") {
		chemical := NewChemical(spec)
		chemicals[chemical.Name] = chemical
	}

	return chemicals
}

func (chemicals Chemicals) OreNeededForFuel(fuelQty int) int {
	extras := map[string]int{}
	needed := map[string]int{"FUEL": fuelQty}
	oreNeeded := 0
	for {
		for chemical, need := range needed {
			delete(needed, chemical)
			if extras[chemical] >= need {
				extras[chemical] -= need
				continue
			} else {
				need -= extras[chemical]
				extras[chemical] = 0
			}

			neededForThis, extra := chemicals[chemical].Make(need)
			for chemical, qty := range neededForThis {
				if chemical == "ORE" {
					oreNeeded += qty
				} else {
					needed[chemical] += qty
				}
			}
			extras[chemical] += extra
		}
		if len(needed) == 0 {
			break
		}
	}
	return oreNeeded
}
