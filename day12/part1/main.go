package main

import "github.com/kitchen/aoc2019/day12"
import "fmt"

func main() {
	moons := day12.Moons{
		day12.NewMoonOnlyPosition(9, 13, -8),
		day12.NewMoonOnlyPosition(-3, 16, -17),
		day12.NewMoonOnlyPosition(-4, 11, -10),
		day12.NewMoonOnlyPosition(0, -2, -2),
	}

	moons = day12.Tick(moons, 1000)
	fmt.Printf("total energy: %v\n", moons.TotalEnergy())
}
