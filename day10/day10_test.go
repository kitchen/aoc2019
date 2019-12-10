package day10

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/suite"
)

type day10Suite struct {
	suite.Suite
}

func (suite *day10Suite) SetupTest() {
}

func (suite *day10Suite) TestParts() {
	p00 := NewPoint(0, 0)
	p10 := NewPoint(1, 0)
	p20 := NewPoint(2, 0)
	p11 := NewPoint(1, 1)
	p22 := NewPoint(2, 2)
	p01 := NewPoint(0, 1)
	p02 := NewPoint(0, 2)

	suite.Equal(1, len(p00.Between(p20)))
	suite.Equal(1, len(p00.Between(p02)))
	suite.Equal(1, len(p00.Between(p22)))

	suite.Equal(Points{p10}, p00.Between(p20))
	suite.Equal(Points{p01}, p00.Between(p02))
	suite.Equal(Points{p11}, p00.Between(p22))

	ps1 := Points{p00, p01, p02, p22}
	ps2 := Points{p00, p22, p11, p20}
	suite.ElementsMatch(Points{p00, p22}, ps1.Union(ps2))

	points := Points{p22, p01, p02}
	suite.Equal(Points{p01, p02, p22}, points.SortByDistance(p00))
}

func (suite *day10Suite) TestExample1() {
	// really don't like golang multiline strings :crysad:
	exampleMap := `.#..#
.....
#####
....#
...##`
	grid := NewGrid(exampleMap)
	suite.Equal(10, len(grid.Points))

	suite.Equal(7, len(grid.VisibleFrom(NewPoint(1, 0))))

	bestPoint, bestVisible := grid.BestVisibility()
	suite.Equal(NewPoint(3, 4), bestPoint)
	suite.Equal(8, bestVisible)

}

func (suite *day10Suite) TestExample2() {
	exampleMap := `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`
	grid := NewGrid(exampleMap)
	bestPoint, bestVisible := grid.BestVisibility()
	point58 := NewPoint(5, 8)
	// point18 := NewPoint(1, 8)
	suite.Equal(point58, bestPoint)
	suite.Equal(33, bestVisible)

	suite.Equal(33, len(grid.VisibleFrom(point58)))

}

func (suite *day10Suite) TestExample3() {
	exampleMap := `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`
	grid := NewGrid(exampleMap)
	bestPoint, bestVisible := grid.BestVisibility()
	suite.Equal(NewPoint(1, 2), bestPoint)
	suite.Equal(35, bestVisible)
}

func (suite *day10Suite) TestExample4() {
	exampleMap := `.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`

	grid := NewGrid(exampleMap)
	bestPoint, bestVisible := grid.BestVisibility()
	suite.Equal(NewPoint(6, 3), bestPoint)
	suite.Equal(41, bestVisible)

}

func (suite *day10Suite) TestExample5() {
	exampleMap := `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`

	grid := NewGrid(exampleMap)
	bestPoint, bestVisible := grid.BestVisibility()
	suite.Equal(NewPoint(11, 13), bestPoint)
	suite.Equal(210, bestVisible)
}

func (suite *day10Suite) TestBlockers() {
	exampleMap := `#.........
...A......
...B..a...
.EDCG....a
..F.c.b...
.....c....
..efd.c.gb
.......c..
....f...c.
...e..d..c`
	re, _ := regexp.Compile("[^\\.\\n]")
	exampleMap = re.ReplaceAllLiteralString(exampleMap, "#")

	grid := NewGrid(exampleMap)
	point00 := NewPoint(0, 0)
	visible := grid.VisibleFrom(point00)
	suite.Equal(7, len(visible))

	// fmt.Printf("visibility map from 0,0\n%v", grid.VisibilityMap(point00))

}

func (suite *day10Suite) TestCandidates() {
	p1 := NewPoint(8, 2)
	p2 := NewPoint(2, 8)
	filler1 := NewPoint(20, 20)
	grid := Grid{Points: Points{p1, p2, filler1}}

	suite.Equal(3, len(grid.Points))

	// fmt.Printf("visibility map from %v\n%v", p1, grid.VisibilityMap(p1))
	// fmt.Printf("candidate map from %v -> %v\n%v", p1, p2, grid.CandidateMap(p1, p2))
}

func (suite *day10Suite) TestZapper() {
	exampleMap := `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`

	grid := NewGrid(exampleMap)
	source := NewPoint(11, 13)

	zapped := grid.Zap(source, 200)
	suite.Equal(NewPoint(11, 12), zapped[0], "first zapped")
	suite.Equal(NewPoint(12, 1), zapped[1], "second zapped")
	suite.Equal(NewPoint(12, 2), zapped[2], "third zapped")
	suite.Equal(NewPoint(12, 8), zapped[9], "tenth zapped")
	suite.Equal(NewPoint(8, 2), zapped[199], "200th zapped")

}

func TestDay10Suite(t *testing.T) {
	suite.Run(t, new(day10Suite))
}
