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

	findX := []int{9, -3, -4, 0}
	prevX := 0
	findY := []int{13, 16, 11, -2}
	prevY := 0
	findZ := []int{-8, -17, -10, -2}
	prevZ := 0

	iteration := 1
	for {
		moons = day12.Tick(moons, 1)

		curX := []int{moons[0].Position.X, moons[1].Position.X, moons[2].Position.X, moons[3].Position.X}
		if testEq(curX, findX) {
			//fmt.Printf("X repeats at %v, cur - prev: %v\n", iteration, iteration-prevX)
			//prevX = iteration
		}
		curY := []int{moons[0].Position.Y, moons[1].Position.Y, moons[2].Position.Y, moons[3].Position.Y}
		if testEq(curY, findY) {
			fmt.Printf("Y repeats at %v, cur - prev: %v\n", iteration, iteration-prevY)
			prevY = iteration
		}
		curZ := []int{moons[0].Position.Z, moons[1].Position.Z, moons[2].Position.Z, moons[3].Position.Z}
		if testEq(curZ, findZ) {
			//fmt.Printf("Z repeats at %v, cur - prev: %v\n", iteration, iteration-prevZ)
			//prevZ = iteration
		}
		iteration += 1
	}
	fmt.Println(prevX, prevY, prevZ)
}

// thanks stackoverflow ... I'm lazy so I just yoinked this from here:
// https://stackoverflow.com/questions/15311969/checking-the-equality-of-two-slices
// really wish golang slices were ==able
func testEq(a, b []int) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// X repeats at 79119599
// X repeats at 79119600
// X repeats at 79232627
// X repeats at 79232628
// X repeats at 79345655
// X repeats at 79345656
// X repeats at 79458683
// X repeats at 79458684
// 113027
// X repeats at 79571711
// X repeats at 79571712
// 113027
// X repeats at 79684739
// X repeats at 79684740
// X repeats at 79797767
// X repeats at 79797768
// X repeats at 79910795
// X repeats at 79910796
// X repeats at 80023823

//
// Y repeats at 16444593
// Y repeats at 16444594
// Y repeats at 16676207
// Y repeats at 16676208
// Y repeats at 16907821
// Y repeats at 16907822
// Y repeats at 17139435
// Y repeats at 17139436
// 231613
// Y repeats at 17371049
// Y repeats at 17371050
// 231613
// Y repeats at 17602663
// Y repeats at 17602664
// Y repeats at 17834277
// Y repeats at 17834278
// Y repeats at 18065891
// Y repeats at 18065892
// Y repeats at 18297505
// Y repeats at 18297506
// Y repeats at 18529119
// Y repeats at 18529120
// Y repeats at 18760733

// Z repeats at 7475735
// Z repeats at 7475736
// Z repeats at 7584079
// Z repeats at 7584080
// Z repeats at 7692423
// Z repeats at 7692424
// Z repeats at 7800767
// Z repeats at 7800768
// Z repeats at 7909111
// Z repeats at 7909112
// 108343
// Z repeats at 8017455
// Z repeats at 8017456
// Z repeats at 8125799
// Z repeats at 8125800
// Z repeats at 8234143
// Z repeats at 8234144
// Z repeats at 8342487
// Z repeats at 8342488
