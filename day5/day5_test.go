package day5

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type day5Suite struct {
	suite.Suite
}

func (suite *day5Suite) SetupTest() {
}

func (suite *day5Suite) TestDay2Stuff() {
	program := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	final, _, _ := Run(program, 0, 0)
	suite.Equal(3500, final[0])

	program = []int{1, 0, 0, 0, 99}
	final, _, _ = Run(program, 0, 0)
	suite.Equal([]int{2, 0, 0, 0, 99}, final)

	program = []int{2, 3, 0, 3, 99}
	final, _, _ = Run(program, 0, 0)
	suite.Equal([]int{2, 3, 0, 6, 99}, final)

	program = []int{2, 4, 4, 5, 99, 0}
	final, _, _ = Run(program, 0, 0)
	suite.Equal([]int{2, 4, 4, 5, 99, 9801}, final)

	program = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	final, _, _ = Run(program, 0, 0)
	suite.Equal([]int{30, 1, 1, 4, 2, 5, 6, 0, 99}, final)
}

func (suite *day5Suite) TestModes() {
	program := []int{3, 0, 4, 0, 99}
	_, output, _ := Run(program, 0, 420)
	suite.Equal(420, output)

	program = []int{1002, 4, 3, 4, 33}
	_, _, err := Run(program, 0, 0)
	// if it even exits it worked
	suite.NoError(err)
}

func (suite *day5Suite) TestComparators() {
	program := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	_, output, _ := Run(program, 0, 8)
	suite.Equal(1, output)
	_, output, _ = Run(program, 0, 42)
	suite.Equal(0, output)

	program = []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	_, output, _ = Run(program, 0, 8)
	suite.Equal(1, output)
	_, output, _ = Run(program, 0, 42)
	suite.Equal(0, output)

	program = []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	_, output, _ = Run(program, 0, 7)
	suite.Equal(1, output)
	_, output, _ = Run(program, 0, 8)
	suite.Equal(0, output)
	_, output, _ = Run(program, 0, 42)
	suite.Equal(0, output)

	program = []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	_, output, _ = Run(program, 0, 7)
	suite.Equal(1, output)
	_, output, _ = Run(program, 0, 8)
	suite.Equal(0, output)
	_, output, _ = Run(program, 0, 42)
	suite.Equal(0, output)

	program = []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	_, output, _ = Run(program, 0, 0)
	suite.Equal(0, output)

	program = []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	_, output, _ = Run(program, 0, 2)
	suite.Equal(1, output)

	program = []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	_, output, _ = Run(program, 0, -1)
	suite.Equal(1, output)

	program = []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	_, output, _ = Run(program, 0, 4)
	suite.Equal(1, output)

	program = []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	_, output, _ = Run(program, 0, 0)
	suite.Equal(0, output)

	program = []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	_, output, _ = Run(program, 0, 1)
	suite.Equal(1, output)

	program = []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	_, output, _ = Run(program, 0, -42)
	suite.Equal(1, output)

	program = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	_, output, _ = Run(program, 0, 0)
	suite.Equal(999, output)

	program = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	_, output, _ = Run(program, 0, 8)
	suite.Equal(1000, output)

	program = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	_, output, _ = Run(program, 0, 3214124)
	suite.Equal(1001, output)
}

func (suite *day5Suite) TestArg() {
	ops := []int{0, 4, 4, 4, 42}
	suite.Equal(42, Arg(ops, 0, 1))
	suite.Equal(42, Arg(ops, 0, 2))
	suite.Equal(42, Arg(ops, 0, 3))

	ops = []int{11100, 23, 34, 45, 42}
	suite.Equal(23, Arg(ops, 0, 1))
	suite.Equal(34, Arg(ops, 0, 2))
	suite.Equal(45, Arg(ops, 0, 3))
}

func TestDay5Suite(t *testing.T) {
	suite.Run(t, new(day5Suite))
}
