package day1

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type day1Suite struct {
	suite.Suite
}

func (suite *day1Suite) SetupTest() {
}

func (suite *day1Suite) TestFuelForMass() {
	suite.Equal(2, FuelForMass(12))
	suite.Equal(2, FuelForMass(14))
	suite.Equal(654, FuelForMass(1969))
	suite.Equal(33583, FuelForMass(100756))
}

func (suite *day1Suite) TestFuelForMassAndFuel() {
	suite.Equal(2, FuelForMassAndFuel(14))
	suite.Equal(966, FuelForMassAndFuel(1969))
	suite.Equal(50346, FuelForMassAndFuel(100756))
}

func TestDay1Suite(t *testing.T) {
	suite.Run(t, new(day1Suite))
}
