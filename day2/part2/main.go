package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kitchen/aoc2019/day2"
)

func main() {
	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	programStrings := strings.Split(strings.Trim(string(buf), "\n"), ",")

	program := make([]int, len(programStrings))
	for i, op := range programStrings {
		program[i], err = strconv.Atoi(op)
		if err != nil {
			log.Fatalf("unable to convert string to int: %v", err)
		}
	}

	// argh. need to read the instructions better :grumble:
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			curProgram := make([]int, len(program))
			copy(curProgram, program)
			curProgram[1] = noun
			curProgram[2] = verb

			output, err := day2.Run(curProgram, 0)
			if err != nil {
				fmt.Printf("noun: %v, verb: %v, error: %v\n", noun, verb, err)
			} else {
				fmt.Printf("noun: %v, verb: %v, 0th: %v\n", noun, verb, output[0])
				if output[0] == 19690720 {
					fmt.Printf("found our match, result: %v\n", 100*noun+verb)
					return
				}
			}

		}
	}

}
