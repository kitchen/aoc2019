package day16

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type day16Suite struct {
	suite.Suite
}

func (suite *day16Suite) SetupTest() {
}

func (suite *day16Suite) TestDigits() {
	suite.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8}, Digits("12345678"))
}

func (suite *day16Suite) TestWeirdlyDescribedDifficultPatternBuilderThing() {
	basePattern := []int{0, 1, 0, -1}

	suite.Equal([]int{1, 0, -1, 0, 1, 0, -1, 0}, patternForDigitAsArray(basePattern, 8, 1))
	suite.Equal([]int{0, 1, 1, 0, 0, -1, -1, 0}, patternForDigitAsArray(basePattern, 8, 2))
	suite.Equal([]int{0, 0, 1, 1, 1, 0, 0, 0}, patternForDigitAsArray(basePattern, 8, 3))
	suite.Equal([]int{0, 0, 0, 1, 1, 1, 1, 0}, patternForDigitAsArray(basePattern, 8, 4))
	suite.Equal([]int{0, 0, 0, 0, 1, 1, 1, 1}, patternForDigitAsArray(basePattern, 8, 5))
	suite.Equal([]int{0, 0, 0, 0, 0, 1, 1, 1}, patternForDigitAsArray(basePattern, 8, 6))
	suite.Equal([]int{0, 0, 0, 0, 0, 0, 1, 1}, patternForDigitAsArray(basePattern, 8, 7))
	suite.Equal([]int{0, 0, 0, 0, 0, 0, 0, 1}, patternForDigitAsArray(basePattern, 8, 8))
}

func (suite *day16Suite) TestExample1() {
	basePattern := []int{0, 1, 0, -1}
	inputDigits := Digits("12345678")

	inputDigits = FFT(basePattern, inputDigits, 1)
	suite.Equal(Digits("48226158"), inputDigits)
	inputDigits = FFT(basePattern, inputDigits, 1)
	suite.Equal(Digits("34040438"), inputDigits)
	inputDigits = FFT(basePattern, inputDigits, 1)
	suite.Equal(Digits("03415518"), inputDigits)
	inputDigits = FFT(basePattern, inputDigits, 1)
	suite.Equal(Digits("01029498"), inputDigits)
}

func (suite *day16Suite) TestExample2() {
	basePattern := []int{0, 1, 0, -1}
	inputDigits := Digits("19617804207202209144916044189917")

	for i := 1; i <= 100; i++ {
		inputDigits = FFT(basePattern, inputDigits, 1)
	}

	suite.Equal(Digits("73745418"), inputDigits[0:8])
}

func (suite *day16Suite) TestExample3() {
	basePattern := []int{0, 1, 0, -1}
	inputDigits := Digits("69317163492948606335995924319873")

	for i := 1; i <= 100; i++ {
		inputDigits = FFT(basePattern, inputDigits, 1)
	}

	suite.Equal(Digits("52432133"), inputDigits[0:8])

}

func (suite *day16Suite) TestPart1() {
	basePattern := []int{0, 1, 0, -1}
	inputDigits := Digits("59781998462438675006185496762485925436970503472751174459080326994618036736403094024111488348676644802419244196591075975610084280308059415695059918368911890852851760032000543205724091764633390765212561307082338287866715489545069566330303873914343745198297391838950197434577938472242535458546669655890258618400619467693925185601880453581947475741536786956920286681271937042394272034410161080365044440682830248774547018223347551308590698989219880430394446893636437913072055636558787182933357009123440661477321673973877875974028654688639313502382365854245311641198762520478010015968789202270746880399268251176490599427469385384364675153461448007234636949")

	for i := 1; i <= 100; i++ {
		inputDigits = FFT(basePattern, inputDigits, 1)
	}

	suite.Equal(Digits("23135243"), inputDigits[0:8])

}

func TestDay16Suite(t *testing.T) {
	suite.Run(t, new(day16Suite))
}
