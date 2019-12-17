package day15

import (
	"fmt"
	"log"
)

const (
	North = iota + 1
	South
	West
	East
)

const (
	Wall = iota
	OK
	Oxygen
)

type Droid struct {
	Computer *Computer
	Screen   *Screen
	Position Pixel
}

func NewDroid(program []int) *Droid {
	inputChan := make(chan int)
	outputChan := make(chan int)
	promptChan := make(chan bool)
	doneChan := make(chan bool)
	computer := NewComputer(program, 0, 0, inputChan, promptChan, outputChan, doneChan)
	screen := NewScreen()
	position := NewPixel(0, 0)
	return &Droid{Computer: computer, Screen: screen, Position: position}
}

func (droid *Droid) DFS(path []Pixel) bool {
	currentPixel := droid.Screen.Pixels[droid.Position]
	if currentPixel == OxygenType {
		fmt.Printf("length of path to oxygen: %v\n", len(path))
		return true
	}
	for direction := 1; direction <= 4; direction++ {
		nextType := droid.Peek(direction)
		if nextType == UnexploredType {
			droid.Go(direction)
			currentPath := make([]Pixel, len(path))
			copy(currentPath, path)
			currentPath = append(currentPath, droid.Position)
			ret := droid.DFS(currentPath)
			if ret {
				return true
			}
			droid.GoBack(direction)
		}
	}
	return false
}

func (droid *Droid) DFSForFarthest(path []Pixel) int {
	maxLen := len(path)
	droid.Screen.Pixels[droid.Position] = OxygenType
	fmt.Printf("droid is at %v\n", droid.Position)
	fmt.Printf("%v", droid.Screen.Draw(droid.Position))
	for direction := 1; direction <= 4; direction++ {
		nextType := droid.Peek(direction)
		if nextType == UnexploredType || nextType == EmptyType {
			droid.Go(direction)
			currentPath := make([]Pixel, len(path))
			copy(currentPath, path)
			currentPath = append(currentPath, droid.Position)
			ret := droid.DFSForFarthest(currentPath)
			droid.GoBack(direction)
			if ret > maxLen {
				maxLen = ret
			}
		}
	}
	return maxLen
}

func (droid *Droid) NextPixel(direction int) Pixel {
	switch direction {
	case North:
		return NewPixel(droid.Position.X, droid.Position.Y-1)
	case South:
		return NewPixel(droid.Position.X, droid.Position.Y+1)
	case East:
		return NewPixel(droid.Position.X+1, droid.Position.Y)
	case West:
		return NewPixel(droid.Position.X-1, droid.Position.Y)
	}
	log.Fatal("invalid direction")
	return Pixel{}
}

func (droid *Droid) Go(direction int) int {
	<-droid.Computer.PromptForInput
	droid.Computer.Input <- direction
	what := <-droid.Computer.Output

	var candidatePosition = droid.NextPixel(direction)
	if _, exists := droid.Screen.Pixels[candidatePosition]; !exists {
		droid.Screen.Pixels[candidatePosition] = PixelType(what + 1)
	}
	if what != Wall {
		droid.Position = candidatePosition
	}
	return what
}

func (droid *Droid) GoBack(direction int) {
	var back int
	switch direction {
	case North:
		back = South
	case South:
		back = North
	case East:
		back = West
	case West:
		back = East
	}
	_ = droid.Go(back)
}

func (droid *Droid) Peek(direction int) PixelType {
	pixel := droid.NextPixel(direction)

	pixelType, exists := droid.Screen.Pixels[pixel]
	if exists {
		return pixelType
	}
	what := droid.Go(direction)
	if what != Wall {
		droid.GoBack(direction)
		return UnexploredType
	}

	return WallType
}
