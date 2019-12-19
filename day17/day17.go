package day17

import "fmt"

const (
	Scaffold     = 35
	OpenSpace    = 46
	NewLine      = 10
	RobotUp      = 94
	RobotRight   = 62
	RobotDown    = 118
	RobotLeft    = 60
	RobotOops    = 88
	Intersection = 79
)

var PixelChars = map[PixelType]string{
	Scaffold:     "#",
	OpenSpace:    ".",
	RobotUp:      "^",
	RobotRight:   ">",
	RobotDown:    "v",
	RobotLeft:    "<",
	RobotOops:    "X",
	Intersection: "O",
}

var PixelTypes = map[string]PixelType{
	"#": Scaffold,
	".": OpenSpace,
	"^": RobotUp,
	">": RobotRight,
	"v": RobotDown,
	"<": RobotLeft,
	"X": RobotOops,
	"O": Intersection,
}

type ASCII struct {
	Computer *Computer
	Screen   *Screen
}

func NewASCII() *ASCII {
	inputChan := make(chan int)
	promptChan := make(chan bool)
	outputChan := make(chan int)
	doneChan := make(chan bool)

	computer := NewComputer(Program, 0, 0, inputChan, promptChan, outputChan, doneChan)
	screen := NewScreen()

	return &ASCII{Computer: computer, Screen: screen}
}

func (ascii *ASCII) Run() map[Pixel]bool {
	go ascii.Computer.Run()
	y := 0
	x := 0
	scaffolds := map[Pixel]bool{}
	for pixel := range ascii.Computer.Output {
		if pixel == NewLine {
			y += 1
			x = 0
			continue
		} else {
			if y == 0 || y == 1 {
				fmt.Printf("%v %v %y\n", x, y, PixelType(pixel))
			}
			ascii.Screen.Pixels[NewPixel(x, y)] = PixelType(pixel)
		}

		if pixel == Scaffold {
			scaffolds[NewPixel(x, y)] = true
		}
		x += 1
	}
	<-ascii.Computer.Done

	return scaffolds
}
