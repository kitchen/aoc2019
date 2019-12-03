package day3

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Path []Point

func (p Point) DistanceFromCenter() int {
	return p.DistanceFrom(Point{x: 0, y: 0})
}

func (p1 Point) DistanceFrom(p2 Point) int {
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
}

func NewPath(pathSpec string) Path {
	x, y := 0, 0
	path := Path{}
	segments := strings.Split(pathSpec, ",")
	for _, segment := range segments {
		direction := segment[:1]
		count, err := strconv.Atoi(segment[1:])
		if err != nil {
			log.Fatalf("couldn't convert string to int: %v", err)
		}

		for i := 1; i <= count; i++ {
			switch direction {
			case "U":
				y++
			case "D":
				y--
			case "R":
				x++
			case "L":
				x--
			default:
				log.Fatalf("invalid direction: %v", direction)
			}
			path = append(path, Point{x: x, y: y})
		}
	}

	return path
}

func (a Path) Union(b Path) []Point {
	aPoints := make(map[Point]bool, len(a))
	for _, point := range a {
		aPoints[point] = true
	}

	common := []Point{}
	for _, point := range b {
		if aPoints[point] {
			common = append(common, point)
		}
	}
	return common
}

func (path Path) StepsTo(point Point) int {
	for steps, checkPoint := range path {
		if checkPoint == point {
			return steps + 1
		}
	}

	log.Fatalf("point isn't in the path")
	return -1
}

func LeastStepsToIntersection(path1 Path, path2 Path) int {
	common := path1.Union(path2)
	var leastSteps int
	for i, point := range common {
		steps := path1.StepsTo(point) + path2.StepsTo(point)
		if i == 0 || steps < leastSteps {
			leastSteps = steps
		}
	}
	return leastSteps
}

func ClosestPoint(points []Point) Point {
	var closest Point
	for i, point := range points {
		if i == 0 || point.DistanceFromCenter() < closest.DistanceFromCenter() {
			closest = point
		}
	}

	return closest
}
