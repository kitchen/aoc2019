package day3

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type day3Suite struct {
	suite.Suite
}

func (suite *day3Suite) SetupTest() {
}

func (suite *day3Suite) TestNewPath() {
	path1Spec := "R8,U5,L5,D3"
	path1 := NewPath(path1Spec)
	suite.Equal(21, len(path1))

	path2Spec := "U7,R6,D4,L4"
	path2 := NewPath(path2Spec)
	suite.Equal(21, len(path2))

	common := path1.Union(path2)
	suite.Equal(2, len(common))

	point := ClosestPoint(common)
	suite.Equal(6, point.DistanceFromCenter())

	path1 = NewPath("R75,D30,R83,U83,L12,D49,R71,U7,L72")
	path2 = NewPath("U62,R66,U55,R34,D71,R55,D58,R83")
	common = path1.Union(path2)
	point = ClosestPoint(common)
	suite.Equal(159, point.DistanceFromCenter())

	path1 = NewPath("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
	path2 = NewPath("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	common = path1.Union(path2)
	point = ClosestPoint(common)
	suite.Equal(135, point.DistanceFromCenter())

}

func (suite *day3Suite) TestPart2() {
	path1 := NewPath("R75,D30,R83,U83,L12,D49,R71,U7,L72")
	path2 := NewPath("U62,R66,U55,R34,D71,R55,D58,R83")
	suite.Equal(610, LeastStepsToIntersection(path1, path2))

	path1 = NewPath("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
	path2 = NewPath("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	suite.Equal(410, LeastStepsToIntersection(path1, path2))

}

func TestDay2Suite(t *testing.T) {
	suite.Run(t, new(day3Suite))
}
