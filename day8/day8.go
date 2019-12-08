package day8

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/mgutz/ansi"
)

type Picture struct {
	Layers []*Layer
	Height int
	Width  int
}

type Layer struct {
	Pixels [][]int
}

func NewPicture(width int, height int, data string) *Picture {
	layerSize := height * width
	layers := []*Layer{}

	for i := 0; i < len(data); i += layerSize {
		layer := NewLayer(width, height, data[i:i+layerSize])
		layers = append(layers, layer)
	}
	return &Picture{Width: width, Height: height, Layers: layers}
}

func (p *Picture) Stacked() [][]int {
	stacked := make([][]int, p.Height)
	for row := 0; row < p.Height; row++ {
		fmt.Printf("row %v\n", row)
		stacked[row] = make([]int, p.Width)
		for col := 0; col < p.Width; col++ {
			color := 0
			fmt.Printf("\tcol %v\n", col)
			for i := len(p.Layers) - 1; i >= 0; i-- {
				layer := p.Layers[i]
				layerColor := layer.Pixels[row][col]
				fmt.Printf("\t\tlayer %v -> %v\n", i, layerColor)
				if layerColor != 2 {
					color = layerColor
				}
			}
			stacked[row][col] = color
			fmt.Printf("\t\t\tfinal: %v\n", color)
		}
	}

	return stacked
}

func (p *Picture) PrintStacked() string {
	stacked := p.Stacked()
	output := ""
	for _, row := range stacked {
		for _, col := range row {
			var color string
			if col == 0 {
				color = "black"
			} else {
				color = "white"
			}
			output += ansi.Color("X", color)
		}
		output += "\n"
	}
	return output
}

func NewLayer(width int, height int, data string) *Layer {
	pixels := make([][]int, height)
	for i, numString := range strings.Split(data, "") {
		row := i / width
		col := i % width
		if col == 0 {
			pixels[row] = make([]int, width)
		}
		pixel, err := strconv.Atoi(numString)
		if err != nil {
			log.Fatal("htns")
		}
		pixels[row][col] = pixel
	}

	return &Layer{Pixels: pixels}
}

func (l *Layer) CountNumbers() map[int]int {
	counts := map[int]int{}
	for _, row := range l.Pixels {
		for _, col := range row {
			counts[col]++
		}
	}
	return counts
}
