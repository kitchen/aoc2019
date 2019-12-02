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
	program[1] = 12
	program[2] = 2

	output, _ := day2.Run(program, 0)

	fmt.Printf("output: %v", output)

	fmt.Printf("0th position value: %v\n", output[0])

}
