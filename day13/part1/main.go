package main

import (
	"fmt"

	"github.com/kitchen/aoc2019/day13"
	"github.com/kitchen/aoc2019/day9"
)

func main() {
	inputChan := make(chan int)
	outputChan := make(chan int)
	doneChan := make(chan bool)
	computer := day9.NewComputer(day13.Program, 0, 0, inputChan, outputChan, doneChan)

	go computer.Run()
	screen := day13.NewScreen()

	for {
		x, y, pixelType := <-outputChan, <-outputChan, <-outputChan
		screen.Pixels[day13.NewPixel(x, y)] = day13.PixelType(pixelType)

		exitLoop := false
		select {
		case <-doneChan:
			exitLoop = true
		default:
		}
		if exitLoop {
			break
		}
	}

	counts := screen.Pixels.CountByType()
	fmt.Printf("number of blocks is %v\n", counts[day13.Block])

}
