package main

import (
	"fmt"

	"github.com/kitchen/aoc2019/day16"
)

func main() {

	input := day16.Digits("59781998462438675006185496762485925436970503472751174459080326994618036736403094024111488348676644802419244196591075975610084280308059415695059918368911890852851760032000543205724091764633390765212561307082338287866715489545069566330303873914343745198297391838950197434577938472242535458546669655890258618400619467693925185601880453581947475741536786956920286681271937042394272034410161080365044440682830248774547018223347551308590698989219880430394446893636437913072055636558787182933357009123440661477321673973877875974028654688639313502382365854245311641198762520478010015968789202270746880399268251176490599427469385384364675153461448007234636949")
	basePattern := []int{0, 1, 0, -1}
	offsetDigits := input[0:7]

	originalInputChan := make(chan int)
	inputChan := originalInputChan
	finalOutputChan := make(chan int)

	go func() {
		for i := 1; i <= 1; i++ {
			for _, digit := range input {
				originalInputChan <- digit
			}
		}
	}()

	for i := 1; i <= 100; i++ {
		var outputChan chan int
		if i == 100 {
			outputChan = finalOutputChan
		} else {
			outputChan = make(chan int)
		}

		go day16.FFT(basePattern, len(input)*10000, inputChan, outputChan)
		inputChan = outputChan
	}

	offset := 0
	for i := 0; i < 7; i++ {
		offset *= 10
		offset += offsetDigits[i]
	}

	// for i := 0; i < offset; i++ {
	// 	<-finalOutputChan
	// }

	for i := 0; i < 8; i++ {
		fmt.Printf("final output digit: %v\n", <-finalOutputChan)
	}
}
