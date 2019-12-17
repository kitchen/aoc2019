package day15

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type day15Suite struct {
	suite.Suite
}

func (suite *day15Suite) SetupTest() {
}

func (suite *day15Suite) TestStuff() {
}

func TestDay15Suite(t *testing.T) {
	suite.Run(t, new(day15Suite))
}
