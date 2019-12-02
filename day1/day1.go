package day1

// moduleFuel = math.Floor(moduleMass / 3) - 2

func FuelForMass(mass int) int {
	return (mass / 3) - 2
}

func FuelForMassAndFuel(mass int) int {
	fuel := FuelForMass(mass)
	if fuel <= 0 {
		return 0
	}
	return fuel + FuelForMassAndFuel(fuel)
}
