package main

import (
	"fmt"

	"github.com/kitchen/aoc2019/day13"
)

func main() {
	inputChan := make(chan int)
	promptChan := make(chan bool)
	outputChan := make(chan int)
	doneChan := make(chan bool)

	curProgram := make([]int, len(day13.Program))
	copy(curProgram, day13.Program)
	curProgram[0] = 2

	computer := day13.NewComputer(curProgram, 0, 0, inputChan, promptChan, outputChan, doneChan)

	go computer.Run()
	screen := day13.NewScreen()
	var ballX, paddleX int

	for {
		var x, y, pixelType int
		exitLoop := false

		select {
		case <-promptChan:
			// fmt.Printf(screen.Draw())
			paddleInput := 0
			if ballX < paddleX {
				paddleInput = -1
			} else if ballX > paddleX {
				paddleInput = 1
			}
			// fmt.Printf("paddle input: %v\n", paddleInput)
			inputChan <- paddleInput
		case x = <-outputChan:
			y, pixelType = <-outputChan, <-outputChan
			if x == -1 && y == 0 {
				screen.Score = pixelType
			} else if pixelType == int(day13.Ball) {
				ballX = x
			} else if pixelType == int(day13.Hpaddle) {
				paddleX = x
			}
			screen.Pixels[day13.NewPixel(x, y)] = day13.PixelType(pixelType)
		case <-doneChan:
			exitLoop = true
		}

		if exitLoop {
			break
		}
	}

	fmt.Printf("final screen:\n%v", screen.Draw())

}
