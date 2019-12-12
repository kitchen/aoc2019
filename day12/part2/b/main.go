package main

import "fmt"

func main() {
	xs := []int{113027, 113028}
	ys := []int{231613, 231614}
	zs := []int{108343, 108344}

	lowest := 0
	for _, x := range xs {
		for _, y := range ys {
			for _, z := range zs {
				lcm := LCM(x, y, z)
				fmt.Printf("%v %v %v -> %v\n", x, y, z, lcm)
				if lowest == 0 || lcm < lowest {
					lowest = lcm
				}
			}
		}
	}
	fmt.Printf("lowest: %v\n", lowest)
}

// more shameless codetheft:
// https://play.golang.org/p/SmzvkDjYlb
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
