package day4

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type day4Suite struct {
	suite.Suite
}

func (suite *day4Suite) SetupTest() {
}

func (suite *day4Suite) TestAllIncreasing() {
	suite.True(AllIncreasing(123456))
	suite.True(AllIncreasing(111111))
	suite.True(AllIncreasing(122334))

	suite.False(AllIncreasing(112231))
	suite.False(AllIncreasing(911111))
}

func (suite *day4Suite) TestHasPairedNumbers() {
	suite.True(HasPairedNumbers(111111))
	suite.True(HasPairedNumbers(123345))
	suite.True(HasPairedNumbers(124566))
	suite.True(HasPairedNumbers(994321))
	suite.False(HasPairedNumbers(123456))
	suite.False(HasPairedNumbers(654321))
}

func (suite *day4Suite) TestHasAnEvenPair() {
	// not a well named function since it's just counting digit values and returning true on
	// 1 or even, but whatever
	suite.True(HasAnEvenPair(112233))
	suite.True(HasAnEvenPair(331144))
	suite.True(HasAnEvenPair(111122))
	suite.False(HasAnEvenPair(111111))
	suite.False(HasAnEvenPair(111123))
	suite.False(HasAnEvenPair(111234))
	suite.True(HasAnEvenPair(222334))

	suite.False(HasAnEvenPair(123444))

	suite.False(HasAnEvenPair(122223))
	suite.True(HasAnEvenPair(111223))
}

func TestDay4Suite(t *testing.T) {
	suite.Run(t, new(day4Suite))
}
