package day9

import (
	"fmt"
	"math"
)

type Computer struct {
	memory       []int
	input        chan int
	output       chan int
	done         chan bool
	relativeBase int
	pos          int
	Iterations   int
}

func Run(memory []int, pos int, input []int) ([]int, []int, error) {
	inputChan := make(chan int)
	outputChan := make(chan int)
	doneChan := make(chan bool)
	var output []int
	computer := NewComputer(memory, pos, 0, inputChan, outputChan, doneChan)

	go computer.Run()

	// fmt.Printf("pushing input\n")
	for _, inputValue := range input {
		inputChan <- inputValue
	}
	close(inputChan)

	// fmt.Printf("pulling output\n")
	for outputValue := range outputChan {
		// fmt.Printf("pulled output\n")
		output = append(output, outputValue)
		// fmt.Printf("done pulled output\n")
	}

	// fmt.Printf("waiting for done\n")
	<-doneChan

	// fmt.Printf("done!\n")
	return computer.memory, output, nil
}

func NewComputer(initMemory []int, initPos int, initRelative int, input chan int, output chan int, done chan bool) *Computer {
	memory := make([]int, 100000)
	copy(memory, initMemory)
	return &Computer{memory: memory, input: input, output: output, pos: initPos, relativeBase: initRelative, done: done, Iterations: 0}
}

func (c *Computer) Run() error {
	c.Iterations++
	// fmt.Printf("program pos: %v\n", c.pos)
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
		c.memory[c.ArgPos(1)] = <-c.input
		// fmt.Printf("%v = input (%v)-> %v\n", c.memory[c.pos:c.pos+2], c.Arg(1), c.ArgPos(1))
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
		// fmt.Println("halting!")
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
