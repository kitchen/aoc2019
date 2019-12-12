package day12

import "math"
import "fmt"

var pairs = [][]int{
	{0, 1},
	{0, 2},
	{0, 3},
	{1, 2},
	{1, 3},
	{2, 3},
}

type Position struct {
	X int
	Y int
	Z int
}

func NewPosition(x, y, z int) Position {
	return Position{X: x, Y: y, Z: z}
}

type Velocity struct {
	X int
	Y int
	Z int
}

func NewVelocity(x, y, z int) Velocity {
	return Velocity{X: x, Y: y, Z: z}
}

type Moon struct {
	Name     string
	Position Position
	Velocity Velocity
}

func NewMoonOnlyPosition(x int, y int, z int) *Moon {
	return &Moon{Position: Position{X: x, Y: y, Z: z}, Velocity: Velocity{X: 0, Y: 0, Z: 0}}
}

func (moon *Moon) GetPulledBy(puller *Moon) {
	switch {
	case moon.Position.X < puller.Position.X:
		moon.Velocity.X++
	case moon.Position.X > puller.Position.X:
		moon.Velocity.X--
	}

	switch {
	case moon.Position.Y < puller.Position.Y:
		moon.Velocity.Y++
	case moon.Position.Y > puller.Position.Y:
		moon.Velocity.Y--
	}

	switch {
	case moon.Position.Z < puller.Position.Z:
		moon.Velocity.Z++
	case moon.Position.Z > puller.Position.Z:
		moon.Velocity.Z--
	}
}

func (moon *Moon) Move() {
	moon.Position.X += moon.Velocity.X
	moon.Position.Y += moon.Velocity.Y
	moon.Position.Z += moon.Velocity.Z
}

func (moon *Moon) Energy() (int, int) {
	potential := int(math.Abs(float64(moon.Position.X)) + math.Abs(float64(moon.Position.Y)) + math.Abs(float64(moon.Position.Z)))
	kinetic := int(math.Abs(float64(moon.Velocity.X)) + math.Abs(float64(moon.Velocity.Y)) + math.Abs(float64(moon.Velocity.Z)))
	return potential, kinetic
}

type Moons []*Moon

func (moons Moons) TotalEnergy() int {
	totalEnergy := 0
	for i, moon := range moons {
		p, k := moon.Energy()
		fmt.Printf("moon %v %v P: %v K: %v P*K: %v\n", i, moon, p, k, p*k)
		totalEnergy += p * k
	}
	return totalEnergy
}

func Tick(moons Moons, count int) Moons {
	if count == 0 {
		return moons
	}
	count -= 1

	for _, pair := range pairs {
		moons[pair[0]].GetPulledBy(moons[pair[1]])
		moons[pair[1]].GetPulledBy(moons[pair[0]])
	}

	for _, moon := range moons {
		moon.Move()
	}

	return Tick(moons, count)
}
