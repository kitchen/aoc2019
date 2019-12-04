package main

import (
	"fmt"

	"github.com/kitchen/aoc2019/day4"
)

func main() {
	count := 0
	for i := 359282; i <= 820401; i++ {
		if day4.AllIncreasing(i) && day4.HasPairedNumbers(i) {
			count++
		}
	}
	fmt.Printf("rule meeting numbers: %v", count)
}
