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

func TestDay5Suite(t *testing.T) {
	suite.Run(t, new(day5Suite))
}
