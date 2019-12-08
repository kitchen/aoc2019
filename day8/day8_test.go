package day8

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type day8Suite struct {
	suite.Suite
}

func (suite *day8Suite) SetupTest() {
}

func (suite *day8Suite) TestStuff() {
	picture := NewPicture(3, 2, "123456789012")

	suite.Equal(2, len(picture.Layers))

	picture = NewPicture(2, 2, "0222112222120000")
	suite.Equal(4, len(picture.Layers))

	suite.Equal([][]int{
		{0, 1},
		{1, 0},
	}, picture.Stacked())

}

func TestDay8Suite(t *testing.T) {
	suite.Run(t, new(day8Suite))
}
