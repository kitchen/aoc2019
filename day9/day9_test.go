package day9

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type day9Suite struct {
	suite.Suite
}

func (suite *day9Suite) SetupTest() {
}

// func (suite *day9Suite) TestDay2Stuff() {
// 	program := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
// 	final, _, _ := Run(program, 0, nil)
// 	suite.Equal(3500, final[0])
//
// 	program = []int{1, 0, 0, 0, 99}
// 	final, _, _ = Run(program, 0, nil)
// 	suite.Equal([]int{2, 0, 0, 0, 99}, final)
//
// 	program = []int{2, 3, 0, 3, 99}
// 	final, _, _ = Run(program, 0, nil)
// 	suite.Equal([]int{2, 3, 0, 6, 99}, final)
//
// 	program = []int{2, 4, 4, 5, 99, 0}
// 	final, _, _ = Run(program, 0, nil)
// 	suite.Equal([]int{2, 4, 4, 5, 99, 9801}, final)
//
// 	program = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
// 	final, _, _ = Run(program, 0, nil)
// 	suite.Equal([]int{30, 1, 1, 4, 2, 5, 6, 0, 99}, final)
// }

func (suite *day9Suite) TestModes() {
	program := []int{3, 0, 4, 0, 99}
	_, output, _ := Run(program, 0, []int{420})
	suite.Equal(420, output[0])

	program = []int{1002, 4, 3, 4, 33}
	_, _, err := Run(program, 0, nil)
	// if it even exits it worked
	suite.NoError(err)
}

func (suite *day9Suite) TestComparators() {
	program := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	_, output, _ := Run(program, 0, []int{8})
	suite.Equal(1, output[0])
	_, output, _ = Run(program, 0, []int{42})
	suite.Equal(0, output[0])

	program = []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	_, output, _ = Run(program, 0, []int{8})
	suite.Equal(1, output[0])
	_, output, _ = Run(program, 0, []int{42})
	suite.Equal(0, output[0])

	program = []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	_, output, _ = Run(program, 0, []int{7})
	suite.Equal(1, output[0])
	_, output, _ = Run(program, 0, []int{8})
	suite.Equal(0, output[0])
	_, output, _ = Run(program, 0, []int{42})
	suite.Equal(0, output[0])

	program = []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	_, output, _ = Run(program, 0, []int{7})
	suite.Equal(1, output[0])
	_, output, _ = Run(program, 0, []int{8})
	suite.Equal(0, output[0])
	_, output, _ = Run(program, 0, []int{42})
	suite.Equal(0, output[0])

	program = []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	_, output, _ = Run(program, 0, []int{0})
	suite.Equal(0, output[0])

	program = []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	_, output, _ = Run(program, 0, []int{2})
	suite.Equal(1, output[0])

	program = []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	_, output, _ = Run(program, 0, []int{-1})
	suite.Equal(1, output[0])

	program = []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	_, output, _ = Run(program, 0, []int{4})
	suite.Equal(1, output[0])

	program = []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	_, output, _ = Run(program, 0, []int{0})
	suite.Equal(0, output[0])

	program = []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	_, output, _ = Run(program, 0, []int{1})
	suite.Equal(1, output[0])

	program = []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	_, output, _ = Run(program, 0, []int{-42})
	suite.Equal(1, output[0])

	program = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	_, output, _ = Run(program, 0, []int{0})
	suite.Equal(999, output[0])

	program = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	_, output, _ = Run(program, 0, []int{8})
	suite.Equal(1000, output[0])

	program = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	_, output, _ = Run(program, 0, []int{3214124})
	suite.Equal(1001, output[0])
}

// func (suite *day9Suite) TestArg() {
// 	ops := []int{0, 4, 4, 4, 42}
// 	suite.Equal(42, Arg(ops, 0, 1))
// 	suite.Equal(42, Arg(ops, 0, 2))
// 	suite.Equal(42, Arg(ops, 0, 3))
//
// 	ops = []int{11100, 23, 34, 45, 42}
// 	suite.Equal(23, Arg(ops, 0, 1))
// 	suite.Equal(34, Arg(ops, 0, 2))
// 	suite.Equal(45, Arg(ops, 0, 3))
// }

func (suite *day9Suite) TestExample1() {
	phases := []int{4, 3, 2, 1, 0}
	input := 0
	program := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	for iter, phaseSetting := range phases {
		localProgram := make([]int, len(program))
		copy(localProgram, program)
		fmt.Printf("running program: %v with phaseSetting %v and input %v\n", localProgram, phaseSetting, input)
		_, output, _ := Run(localProgram, 0, []int{phaseSetting, input})
		input = output[0]
		fmt.Printf("iteration %v output: %v\n", iter, input)
	}

	suite.Equal(43210, input)
}

func (suite *day9Suite) TestExample2() {
	phases := []int{0, 1, 2, 3, 4}
	input := 0
	program := []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}
	for _, phaseSetting := range phases {
		localProgram := make([]int, len(program))
		copy(localProgram, program)
		fmt.Printf("running program: %v with phaseSetting %v and input %v\n", localProgram, phaseSetting, input)
		_, output, _ := Run(localProgram, 0, []int{phaseSetting, input})
		input = output[0]
	}

	suite.Equal(54321, input)
}

func (suite *day9Suite) TestExample3() {
	phases := []int{1, 0, 4, 3, 2}
	input := 0
	program := []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}
	for _, phaseSetting := range phases {
		localProgram := make([]int, len(program))
		copy(localProgram, program)
		_, output, _ := Run(localProgram, 0, []int{phaseSetting, input})
		input = output[0]
	}

	suite.Equal(65210, input)
}

func (suite *day9Suite) TestDay9Examples() {
	program := []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	_, output, _ := Run(program, 0, []int{})
	suite.Equal(1219070632396864, output[0])

	program = []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	_, output, _ = Run(program, 0, []int{})
	suite.Equal([]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, output)

	program = []int{104, 1125899906842624, 99}
	_, output, _ = Run(program, 0, []int{})
	suite.Equal(1125899906842624, output[0])
}

func TestDay7Suite(t *testing.T) {
	suite.Run(t, new(day9Suite))
}
