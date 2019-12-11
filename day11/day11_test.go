package day11

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type day11Suite struct {
	suite.Suite
}

func (suite *day11Suite) SetupTest() {
}

func (suite *day11Suite) TestStuff() {

}

func TestDay11Suite(t *testing.T) {
	suite.Run(t, new(day11Suite))
}
