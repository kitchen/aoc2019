package day13

import (
	"fmt"
	"math"
)

type PixelType int

const (
	Empty PixelType = iota
	Wall
	Block
	Hpaddle
	Ball
)

var PixelChars = map[PixelType]string{
	Empty:   " ",
	Wall:    "#",
	Block:   "X",
	Hpaddle: "_",
	Ball:    ".",
}

type Pixel struct {
	X int
	Y int
}

func NewPixel(x, y int) Pixel {
	return Pixel{X: x, Y: y}
}

type Pixels map[Pixel]PixelType

func (pixels Pixels) CountByType() map[PixelType]int {
	counts := map[PixelType]int{}
	for _, pixelType := range pixels {
		counts[pixelType]++
	}
	return counts
}

type Screen struct {
	Pixels Pixels
	Score  int
}

func NewScreen() *Screen {
	return &Screen{Pixels: Pixels{}, Score: 0}
}

func (screen *Screen) Draw() string {
	var maxX, maxY int
	for pixel, _ := range screen.Pixels {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}

	pixels := ""
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			pixelType := screen.Pixels[NewPixel(x, y)]
			pixels += PixelChars[pixelType]
		}
		pixels += "\n"
	}
	pixels += fmt.Sprintf("score: %v\n", screen.Score)
	return pixels
}

type Computer struct {
	memory         []int
	input          chan int
	promptForInput chan bool
	output         chan int
	done           chan bool
	relativeBase   int
	pos            int
	Iterations     int
}

func NewComputer(initMemory []int, initPos int, initRelative int, input chan int, promptForInput chan bool, output chan int, done chan bool) *Computer {
	memory := make([]int, 100000)
	copy(memory, initMemory)
	return &Computer{memory: memory, input: input, promptForInput: promptForInput, output: output, pos: initPos, relativeBase: initRelative, done: done, Iterations: 0}
}

func (c *Computer) Run() error {
	c.Iterations++
	op := c.memory[c.pos] % 100

	switch op {
	case 1:
		// fmt.Printf("%v = add %v %v -> %v\n", c.memory[c.pos:c.pos+4], c.Arg(1), c.Arg(2), c.ArgPos(3))
		c.memory[c.ArgPos(3)] = c.Arg(1) + c.Arg(2)
		c.pos += 4
	case 2:
		// fmt.Printf("%v = multi %v %v -> %v\n", c.memory[c.pos:c.pos+4], c.Arg(1), c.Arg(2), c.ArgPos(3))
		c.memory[c.ArgPos(3)] = c.Arg(1) * c.Arg(2)
		c.pos += 4
	case 3:
		// fmt.Printf("prompting for input\n")
		c.promptForInput <- true
		c.memory[c.ArgPos(1)] = <-c.input
		// fmt.Printf("%v = input %v -> %v\n", c.memory[c.pos:c.pos+2], c.Arg(1), c.ArgPos(1))
		c.pos += 2
	case 4:
		// fmt.Printf("%v = output %v\n", c.memory[c.pos:c.pos+2], c.Arg(1))
		c.output <- c.Arg(1)
		c.pos += 2
	case 5:
		// fmt.Printf("%v = nzj %v -> %v / %v\n", c.memory[c.pos:c.pos+3], c.Arg(1), c.Arg(2), c.pos+3)
		if c.Arg(1) != 0 {
			c.pos = c.Arg(2)
		} else {
			c.pos += 3
		}
	case 6:
		// fmt.Printf("%v = zj %v -> %v / %v\n", c.memory[c.pos:c.pos+3], c.Arg(1), c.Arg(2), c.pos+3)
		if c.Arg(1) == 0 {
			c.pos = c.Arg(2)
		} else {
			c.pos += 3
		}
	case 7:
		// fmt.Printf("%v = %v lt %v -> %v\n", c.memory[c.pos:c.pos+4], c.Arg(1), c.Arg(2), c.ArgPos(3))
		if c.Arg(1) < c.Arg(2) {
			c.memory[c.ArgPos(3)] = 1
		} else {
			c.memory[c.ArgPos(3)] = 0
		}
		c.pos += 4
	case 8:
		// fmt.Printf("%v = %v eq %v -> %v\n", c.memory[c.pos:c.pos+4], c.Arg(1), c.Arg(2), c.ArgPos(3))
		if c.Arg(1) == c.Arg(2) {
			c.memory[c.ArgPos(3)] = 1
		} else {
			c.memory[c.ArgPos(3)] = 0
		}
		c.pos += 4
	case 9:
		// fmt.Printf("%v = relative + %v -> %v\n", c.memory[c.pos:c.pos+2], c.Arg(1), c.relativeBase+c.Arg(1))
		c.relativeBase += c.Arg(1)
		c.pos += 2
	case 99:
		fmt.Println("halting!")
		close(c.output)
		c.done <- true
		return nil
	default:
		return fmt.Errorf("bad operation")
	}

	return c.Run()
}

func (c *Computer) ArgPos(arg int) int {
	flags := c.memory[c.pos] / 100
	argflag := (flags % int(math.Pow10(arg))) / int(math.Pow10(arg-1))
	if argflag == 1 {
		return c.pos + arg
	} else if argflag == 2 {
		return c.relativeBase + c.memory[c.pos+arg]
	}
	return c.memory[c.pos+arg]
}

func (c *Computer) Arg(arg int) int {
	return c.memory[c.ArgPos(arg)]
}
