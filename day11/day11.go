package day11

import (
	"fmt"

	"github.com/kitchen/aoc2019/day9"
)

type Panel struct {
	x int
	y int
}

type direction int

const (
	up direction = iota
	down
	left
	right
)

type ShipPainter struct {
	computer       *day9.Computer
	Panels         map[Panel]int
	computerInput  chan int
	computerOutput chan int
	Done           chan bool
	position       Panel
	orientation    direction
}

func NewShipPainter(program []int, originColor int) *ShipPainter {
	inputChan := make(chan int)
	outputChan := make(chan int)
	doneChan := make(chan bool)
	computer := day9.NewComputer(program, 0, 0, inputChan, outputChan, doneChan)
	position := Panel{x: 0, y: 0}

	go computer.Run()

	panels := map[Panel]int{}
	panels[position] = originColor

	return &ShipPainter{computer: computer, Panels: panels, computerInput: inputChan, computerOutput: outputChan, Done: doneChan, position: position, orientation: up}
}

// Paint
func (sp *ShipPainter) Paint() bool {
	fmt.Printf("getting ready to paint %v\n", sp.position)
	color, exists := sp.Panels[sp.position]
	if !exists {
		color = 0
	}

	select {
	case sp.computerInput <- color:
	case <-sp.Done:
		return true
	}

	paintColor := <-sp.computerOutput
	fmt.Printf("painting %v -> %v\n", paintColor, sp.position)
	sp.Panels[sp.position] = paintColor
	move := <-sp.computerOutput

	curOr := sp.orientation
	switch move {
	case 0:
		switch sp.orientation {
		case down:
			sp.orientation = right
		case right:
			sp.orientation = up
		case up:
			sp.orientation = left
		case left:
			sp.orientation = down
		}
	case 1:
		switch sp.orientation {
		case down:
			sp.orientation = left
		case right:
			sp.orientation = down
		case up:
			sp.orientation = right
		case left:
			sp.orientation = up
		}
	}
	fmt.Printf("move %v orientation %v -> %v\n", move, curOr, sp.orientation)

	cur := sp.position
	switch sp.orientation {
	case down:
		sp.position = Panel{x: cur.x, y: cur.y + 1}
	case right:
		sp.position = Panel{x: cur.x + 1, y: cur.y}
	case up:
		sp.position = Panel{x: cur.x, y: cur.y - 1}
	case left:
		sp.position = Panel{x: cur.x - 1, y: cur.y}
	}

	fmt.Printf("moved to %v\n", sp.position)
	return false
}
