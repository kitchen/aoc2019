package main

import (
	"fmt"

	"github.com/kitchen/aoc2019/day17"
)

func main() {
	robot := day17.NewASCII()
	scaffolds := robot.Run()

	intersections := []day17.Pixel{}
	overrides := map[day17.Pixel]day17.PixelType{}
	runningAlignmentParameterSum := 0
	for cur, _ := range scaffolds {
		x := cur.X
		y := cur.Y
		left := day17.NewPixel(x-1, y)
		right := day17.NewPixel(x+1, y)
		up := day17.NewPixel(x, y-1)
		down := day17.NewPixel(x, y+1)

		numAround := 0
		for _, pixel := range []day17.Pixel{left, right, up, down} {
			if _, exists := scaffolds[pixel]; exists {
				numAround += 1
			}
		}
		if numAround >= 3 {
			intersections = append(intersections, cur)
			overrides[cur] = day17.Intersection
			runningAlignmentParameterSum += x * y
		}
	}

	overrides[day17.NewPixel(60, 0)] = day17.RobotOops
	overrides[day17.NewPixel(60, 1)] = day17.RobotOops

	fmt.Printf("current screen:\n%v", robot.Screen.Draw(overrides))
	fmt.Printf("intersections: %v\n", intersections)
	fmt.Printf("sum of 'alignment parameters': %v\n", runningAlignmentParameterSum)
}
