package main

import (
	"fmt"

	"github.com/kitchen/aoc2019/day11"
)

func main() {

	program := []int{3, 8, 1005, 8, 318, 1106, 0, 11, 0, 0, 0, 104, 1, 104, 0, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 108, 1, 8, 10, 4, 10, 1002, 8, 1, 28, 1, 107, 14, 10, 1, 107, 18, 10, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 108, 1, 8, 10, 4, 10, 102, 1, 8, 58, 1006, 0, 90, 2, 1006, 20, 10, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1001, 8, 0, 88, 2, 103, 2, 10, 2, 4, 7, 10, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1001, 8, 0, 118, 1, 1009, 14, 10, 1, 1103, 9, 10, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1002, 8, 1, 147, 1006, 0, 59, 1, 104, 4, 10, 2, 106, 18, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 101, 0, 8, 181, 2, 4, 17, 10, 1006, 0, 36, 1, 107, 7, 10, 2, 1008, 0, 10, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 101, 0, 8, 217, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 101, 0, 8, 240, 1006, 0, 64, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1002, 8, 1, 264, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1001, 8, 0, 287, 1, 1104, 15, 10, 1, 102, 8, 10, 1006, 0, 2, 101, 1, 9, 9, 1007, 9, 940, 10, 1005, 10, 15, 99, 109, 640, 104, 0, 104, 1, 21102, 932700857236, 1, 1, 21101, 335, 0, 0, 1106, 0, 439, 21101, 0, 387511792424, 1, 21101, 346, 0, 0, 1106, 0, 439, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 21101, 46372252675, 0, 1, 21102, 393, 1, 0, 1106, 0, 439, 21101, 97806162983, 0, 1, 21102, 404, 1, 0, 1105, 1, 439, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 0, 21102, 1, 825452438376, 1, 21101, 0, 427, 0, 1106, 0, 439, 21102, 709475586836, 1, 1, 21101, 0, 438, 0, 1106, 0, 439, 99, 109, 2, 22101, 0, -1, 1, 21101, 40, 0, 2, 21102, 1, 470, 3, 21102, 1, 460, 0, 1106, 0, 503, 109, -2, 2106, 0, 0, 0, 1, 0, 0, 1, 109, 2, 3, 10, 204, -1, 1001, 465, 466, 481, 4, 0, 1001, 465, 1, 465, 108, 4, 465, 10, 1006, 10, 497, 1101, 0, 0, 465, 109, -2, 2105, 1, 0, 0, 109, 4, 2102, 1, -1, 502, 1207, -3, 0, 10, 1006, 10, 520, 21102, 1, 0, -3, 21202, -3, 1, 1, 21202, -2, 1, 2, 21101, 0, 1, 3, 21101, 0, 539, 0, 1106, 0, 544, 109, -4, 2105, 1, 0, 109, 5, 1207, -3, 1, 10, 1006, 10, 567, 2207, -4, -2, 10, 1006, 10, 567, 22101, 0, -4, -4, 1106, 0, 635, 21202, -4, 1, 1, 21201, -3, -1, 2, 21202, -2, 2, 3, 21102, 586, 1, 0, 1105, 1, 544, 22101, 0, 1, -4, 21102, 1, 1, -1, 2207, -4, -2, 10, 1006, 10, 605, 21102, 0, 1, -1, 22202, -2, -1, -2, 2107, 0, -3, 10, 1006, 10, 627, 22101, 0, -1, 1, 21102, 1, 627, 0, 106, 0, 502, 21202, -2, -1, -2, 22201, -4, -2, -4, 109, -5, 2105, 1, 0}

	painter := day11.NewShipPainter(program, 0)
	for {
		exitLoop := painter.Paint()
		if exitLoop {
			break
		}
	}

	fmt.Printf("number of tiles painted: %v\n", len(painter.Panels))
}
