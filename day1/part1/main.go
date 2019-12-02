package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/kitchen/aoc2019/day1"
)

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("unable to open file %v: %v", os.Args[1], err)
	}

	scanner := bufio.NewScanner(file)

	totalFuel := 0
	for scanner.Scan() {
		line := scanner.Text()
		mass, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("couldn't parse line: %v\n", err)
		} else {
			totalFuel += day1.FuelForMass(mass)
		}
	}

	fmt.Printf("total fuel required: %v\n", totalFuel)
}
