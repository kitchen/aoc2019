package day13

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type day13Suite struct {
	suite.Suite
}

func (suite *day13Suite) SetupTest() {
}

func (suite *day13Suite) TestStuff() {
}

func TestDay13Suite(t *testing.T) {
	suite.Run(t, new(day13Suite))
}
