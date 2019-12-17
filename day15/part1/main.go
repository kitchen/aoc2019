package main

import (
	"fmt"
	"log"

	"github.com/kitchen/aoc2019/day15"
)

func main() {
	droid := day15.NewDroid(day15.Program)
	go func() {
		err := droid.Computer.Run()
		log.Fatal("computer run returned an error: %v\n", err)
	}()
	droid.DFS([]day15.Pixel{})

	var oxygenPixel day15.Pixel
	for pixel, pixelType := range droid.Screen.Pixels {
		if pixelType == day15.OxygenType {
			oxygenPixel = pixel
		}
	}
	fmt.Printf("map looks like:\n%v\n", droid.Screen.Draw(day15.NewPixel(0, 0)))
	fmt.Printf("oxygen is at %v\n", oxygenPixel)
	fmt.Printf("droid is at %v\n", droid.Position)

	// part 2
	maxLen := droid.DFSForFarthest([]day15.Pixel{})
	fmt.Printf("farthest from oxygen is %v\n", maxLen)

}
