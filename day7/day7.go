package day7

import (
	"fmt"
	"math"
)

func Run(ops []int, pos int, input []int) ([]int, []int, error) {

	inputChan := make(chan int)
	outputChan := make(chan int)
	doneChan := make(chan bool)
	var output []int
	go RunWithChannels(ops, pos, inputChan, outputChan, doneChan)

	for _, inputValue := range input {
		inputChan <- inputValue
	}
	close(inputChan)

	for outputValue := range outputChan {
		output = append(output, outputValue)
	}

	<-doneChan

	return ops, output, nil
}

func RunWithChannels(ops []int, pos int, input chan int, output chan int, done chan bool) error {
	// fmt.Printf("program (pos: %v): %v\n", pos, ops)
	op := ops[pos] % 100

	switch op {
	case 1:
		// fmt.Printf("%v = add %v %v -> %v\n", ops[pos:pos+4], Arg(ops, pos, 1), Arg(ops, pos, 2), ops[pos+3])
		ops[ops[pos+3]] = Arg(ops, pos, 1) + Arg(ops, pos, 2)
		pos += 4
	case 2:
		// fmt.Printf("%v = multi %v %v -> %v\n", ops[pos:pos+4], Arg(ops, pos, 1), Arg(ops, pos, 2), ops[pos+3])
		ops[ops[pos+3]] = Arg(ops, pos, 1) * Arg(ops, pos, 2)
		pos += 4
	case 3:
		ops[ops[pos+1]] = <-input
		fmt.Printf("%v = input (%v)-> %v\n", ops[pos:pos+2], ops[ops[pos+1]], ops[pos+1])
		pos += 2
	case 4:
		fmt.Printf("%v = output %v\n", ops[pos:pos+2], Arg(ops, pos, 1))
		output <- Arg(ops, pos, 1)
		pos += 2
	case 5:
		if Arg(ops, pos, 1) != 0 {
			pos = Arg(ops, pos, 2)
		} else {
			pos += 3
		}
	case 6:
		if Arg(ops, pos, 1) == 0 {
			pos = Arg(ops, pos, 2)
		} else {
			pos += 3
		}
	case 7:
		if Arg(ops, pos, 1) < Arg(ops, pos, 2) {
			ops[ops[pos+3]] = 1
		} else {
			ops[ops[pos+3]] = 0
		}
		pos += 4
	case 8:
		if Arg(ops, pos, 1) == Arg(ops, pos, 2) {
			ops[ops[pos+3]] = 1
		} else {
			ops[ops[pos+3]] = 0
		}
		pos += 4
	case 99:
		fmt.Println("halting!")
		close(output)
		done <- true
		return nil
	default:
		return fmt.Errorf("bad operation")
	}

	return RunWithChannels(ops, pos, input, output, done)
}

func Arg(ops []int, pos int, arg int) int {
	flags := ops[pos] / 100
	argflag := (flags % int(math.Pow10(arg))) / int(math.Pow10(arg-1))
	if argflag == 1 {
		return ops[pos+arg]
	}
	return ops[ops[pos+arg]]
}
