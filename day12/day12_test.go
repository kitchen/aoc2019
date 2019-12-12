package day12

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type day12Suite struct {
	suite.Suite
}

func (suite *day12Suite) SetupTest() {
}

func (suite *day12Suite) TestExample1() {

	moons := Moons{
		NewMoonOnlyPosition(-1, 0, 2),
		NewMoonOnlyPosition(2, -10, -7),
		NewMoonOnlyPosition(4, -8, 8),
		NewMoonOnlyPosition(3, 5, -1),
	}

	moons = Tick(moons, 1)
	suite.Equal(NewPosition(2, -1, 1), moons[0].Position)
	suite.Equal(NewVelocity(3, -1, -1), moons[0].Velocity)

	moons = Tick(moons, 9)
	suite.Equal(NewPosition(2, 1, -3), moons[0].Position)
	suite.Equal(NewVelocity(-3, -2, 1), moons[0].Velocity)

	p, k := moons[0].Energy()
	suite.Equal(6, p)
	suite.Equal(6, k)

	p, k = moons[1].Energy()
	suite.Equal(9, p)
	suite.Equal(5, k)

	p, k = moons[2].Energy()
	suite.Equal(10, p)
	suite.Equal(8, k)

	p, k = moons[3].Energy()
	suite.Equal(6, p)
	suite.Equal(3, k)
}

func (suite *day12Suite) TestExample2() {
	moons := Moons{
		// <x=-8, y=-10, z=0>
		NewMoonOnlyPosition(-8, -10, 0),
		// <x=5, y=5, z=10>
		NewMoonOnlyPosition(5, 5, 10),
		// <x=2, y=-7, z=3>
		NewMoonOnlyPosition(2, -7, 3),
		// <x=9, y=-8, z=-3>
		NewMoonOnlyPosition(9, -8, -3),
	}

	moons = Tick(moons, 100)
	suite.Equal(1940, moons.TotalEnergy())

}

func TestDay12Suite(t *testing.T) {
	suite.Run(t, new(day12Suite))
}
