package day17

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type day17Suite struct {
	suite.Suite
}

func (suite *day17Suite) SetupTest() {
}

func (suite *day17Suite) TestStuff() {

}

func TestDay17Suite(t *testing.T) {
	suite.Run(t, new(day17Suite))
}
