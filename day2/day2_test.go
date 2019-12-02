package day2

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type day2Suite struct {
	suite.Suite
}

func (suite *day2Suite) SetupTest() {
}

func (suite *day2Suite) TestProgram() {
	program := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	final, _ := Run(program, 0)
	suite.Equal(3500, final[0])

	program = []int{1, 0, 0, 0, 99}
	final, _ = Run(program, 0)
	suite.Equal([]int{2, 0, 0, 0, 99}, final)

	program = []int{2, 3, 0, 3, 99}
	final, _ = Run(program, 0)
	suite.Equal([]int{2, 3, 0, 6, 99}, final)

	program = []int{2, 4, 4, 5, 99, 0}
	final, _ = Run(program, 0)
	suite.Equal([]int{2, 4, 4, 5, 99, 9801}, final)

	program = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	final, _ = Run(program, 0)
	suite.Equal([]int{30, 1, 1, 4, 2, 5, 6, 0, 99}, final)
}

func TestDay2Suite(t *testing.T) {
	suite.Run(t, new(day2Suite))
}
