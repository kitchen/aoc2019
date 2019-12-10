package day10

import (
	"math"
	"sort"
	"strings"

	"github.com/mgutz/ansi"
)

type Grid struct {
	Points Points
}

func NewGrid(grid string) Grid {
	lines := strings.Split(grid, "\n")
	points := Points{}
	for y, line := range lines {
		linePoints := strings.Split(line, "")
		for x, linePoint := range linePoints {
			if linePoint == "#" {
				points = append(points, NewPoint(x, y))
			}
		}
	}
	return Grid{Points: points}
}

func (g Grid) VisibleFrom(source Point) Points {
	visible := Points{}
	for _, point := range g.Points {
		if point == source {
			continue
		}
		between := source.Between(point)
		if len(between.Union(g.Points)) == 0 {
			visible = append(visible, point)
		}
	}

	return visible
}

func (g Grid) VisibilityMap(source Point) string {
	visible := g.VisibleFrom(source).Map()
	colors := map[Point]string{}
	for _, point := range g.Points {
		if point == source {
			colors[point] = "blue"
		} else if visible[point] {
			colors[point] = "green"
		} else {
			colors[point] = "red"
		}
	}
	return g.DrawMap(colors)
}

func (g Grid) CandidateMap(source Point, to Point) string {
	colors := map[Point]string{}
	colors[source] = "blue"
	colors[to] = "green"

	between := source.Between(to)
	for _, point := range between {
		colors[point] = "red"
	}
	for _, point := range g.Points.Union(between) {
		colors[point] = "pink"
	}

	return g.DrawMap(colors)
}

func (g Grid) DrawMap(colors map[Point]string) string {
	mapString := ""
	maxX := 0
	maxY := 0
	for _, point := range g.Points {
		if point.X > maxX {
			maxX = point.X
		}
		if point.Y > maxY {
			maxY = point.Y
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			point := NewPoint(x, y)
			if color, exists := colors[point]; exists {
				mapString += ansi.Color("X", color)
			} else {
				mapString += ansi.Color(".", "black")
			}
		}
		mapString += "\n"
	}
	return mapString
}

func (g Grid) BestVisibility() (Point, int) {
	var bestPoint Point
	var bestVisible int
	for _, point := range g.Points {
		visible := len(g.VisibleFrom(point))
		if visible > bestVisible {
			bestPoint = point
			bestVisible = visible
		}
	}
	return bestPoint, bestVisible
}

type Points []Point

func (p1 Points) Union(p2 Points) Points {
	p1Map := p1.Map()
	points := Points{}
	for _, point := range p2 {
		if p1Map[point] {
			points = append(points, point)
		}
	}
	return points
}

func (points Points) Map() PointsMap {
	pointsMap := map[Point]bool{}
	for _, point := range points {
		pointsMap[point] = true
	}
	return pointsMap
}

func (points Points) SortByDistance(from Point) Points {
	sort.Slice(points, func(i, j int) bool {
		return points[i].Distance(from) < points[j].Distance(from)
	})
	return points
}

type PointsMap map[Point]bool

func (pm PointsMap) Slice() Points {
	slice := Points{}
	for point, _ := range pm {
		slice = append(slice, point)
	}
	return slice
}

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) Point {
	return Point{X: x, Y: y}
}

func (p1 Point) Between(p2 Point) Points {
	points := Points{p1, p2}
	sort.Slice(points, func(i, j int) bool {
		return points[i].X < points[j].X
	})

	a := points[0]
	b := points[1]
	deltaX := b.X - a.X
	deltaY := b.Y - a.Y
	slope := float64(deltaY) / float64(deltaX)

	between := PointsMap{}

	for x := a.X + 1; x < b.X; x += 1 {
		y := float64(x-a.X)*slope + float64(a.Y)
		if y == float64(int(y)) {
			between[NewPoint(x, int(y))] = true
		}
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].Y < points[j].Y
	})
	a = points[0]
	b = points[1]
	deltaX = b.X - a.X
	deltaY = b.Y - a.Y
	slope = float64(deltaX) / float64(deltaY)
	for y := a.Y + 1; y < b.Y; y += 1 {
		x := float64(y-a.Y)*slope + float64(a.X)
		if x == float64(int(x)) {
			between[NewPoint(int(x), y)] = true
		}
	}

	return between.Slice()
}

func (p1 Point) Distance(p2 Point) int {
	return int(math.Abs(float64(p1.X-p2.X) + math.Abs(float64(p1.Y-p2.Y))))
}
