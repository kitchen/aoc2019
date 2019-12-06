package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/kitchen/aoc2019/day6"
)

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("unable to open file %v: %v", os.Args[1], err)
	}

	scanner := bufio.NewScanner(file)

	dag := day6.NewOrbitalDag()

	for scanner.Scan() {
		line := scanner.Text()
		err := dag.AddOrbit(line)
		if err != nil {
			fmt.Printf("couldn't add orbit: %v\n", err)
		}
	}

	you, err := dag.GetVertex("YOU")
	if err != nil {
		log.Fatal(err)
	}

	san, err := dag.GetVertex("SAN")
	if err != nil {
		log.Fatal(err)
	}

	common := dag.FirstCommonAncestor(san, you)
	distance := dag.DistanceToAncestor(san, common) + dag.DistanceToAncestor(you, common)

	fmt.Printf("total distance: %v\n", distance)
}
