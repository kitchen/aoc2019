package day6

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type day6Suite struct {
	suite.Suite
	dag *OrbitalDAG
}

func (suite *day6Suite) SetupTest() {
	suite.dag = NewOrbitalDag()
}

func (suite *day6Suite) TestStuff() {
	orbits := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}
	for _, orbit := range orbits {
		err := suite.dag.AddOrbit(orbit)
		suite.NoError(err)
	}

	sources := suite.dag.SourceVertices()
	suite.Equal(1, len(sources))
	suite.Equal("COM", sources[0].ID)

	suite.Equal(4, len(suite.dag.SinkVertices()))

	vertex, err := suite.dag.GetVertex("D")
	suite.NoError(err)
	ancestors, err := suite.dag.Ancestors(vertex)
	suite.NoError(err)
	suite.Equal(3, len(ancestors))

	vertex, err = suite.dag.GetVertex("L")
	suite.NoError(err)
	ancestors, err = suite.dag.Ancestors(vertex)
	suite.NoError(err)
	suite.Equal(7, len(ancestors))

	com, err := suite.dag.GetVertex("COM")
	suite.NoError(err)
	suite.Equal(42, suite.dag.Distances(com, 1))
}

func (suite *day6Suite) TestPart2Stuff() {
	orbits := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}

	for _, orbit := range orbits {
		err := suite.dag.AddOrbit(orbit)
		suite.NoError(err)
	}

	san, err := suite.dag.GetVertex("SAN")
	suite.NoError(err)
	you, err := suite.dag.GetVertex("YOU")
	suite.NoError(err)

	ancestors, err := suite.dag.Ancestors(san)
	suite.NoError(err)
	for _, ancestor := range ancestors {
		fmt.Printf("id: %v\n", ancestor.ID)
	}

	d, err := suite.dag.GetVertex("D")
	suite.NoError(err)
	suite.Equal(d, suite.dag.FirstCommonAncestor(you, san))

	suite.Equal(1, suite.dag.DistanceToAncestor(san, d))
	suite.Equal(3, suite.dag.DistanceToAncestor(you, d))
}

func TestDay6Suite(t *testing.T) {
	suite.Run(t, new(day6Suite))
}
