package day16

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func FFT(basePattern []int, numDigits int, inputChan chan int, outputChan chan int) {
	inputDigits := make([]int, numDigits)
	i := 0
	for digit := range inputChan {
		inputDigits[i] = digit
	}
	fmt.Printf("done being fed\n")

	for outputDigitPosition := 0; outputDigitPosition < numDigits; outputDigitPosition++ {
		outputDigit := 0
		patternChan := make(chan int)
		doneChan := make(chan bool)
		go PatternForDigit(basePattern, numDigits, outputDigitPosition+1, patternChan, doneChan)
		<-patternChan
		for _, inputDigit := range inputDigits {
			outputDigit += <-patternChan + inputDigit
		}
		doneChan <- true
		outputDigit = int(math.Abs(float64(outputDigit)))
		outputDigit %= 10
		outputChan <- outputDigit
	}
	close(outputChan)
}

func Digits(input string) []int {
	digitStrings := strings.Split(input, "")
	digits := make([]int, len(digitStrings))
	for i, digitString := range digitStrings {
		digit, err := strconv.Atoi(digitString)
		if err != nil {
			log.Fatalf("couldn't parse digit %v", err)
		}
		digits[i] = digit
	}
	return digits
}

func PatternForDigit(basePattern []int, numDigits int, outputPosition int, outputChan chan int, done chan bool) {
	basePatternDigit := 0
	patternDigits := []int{}
	i := 0
	for {
		for count := 0; count < outputPosition; count++ {
			select {
			case outputChan <- basePattern[basePatternDigit]:
			case <-done:
				return
			}
			patternDigits = append(patternDigits, basePattern[basePatternDigit])
		}
		basePatternDigit += 1
		basePatternDigit %= len(basePattern)
		if i > numDigits {
			break
		}
		i++
	}
}

func patternForDigitAsArray(basePattern []int, numDigits int, whichDigit int) []int {
	doneChan := make(chan bool)
	patternChan := make(chan int)
	output := make([]int, numDigits)
	go PatternForDigit(basePattern, numDigits, whichDigit, patternChan, doneChan)
	<-patternChan
	for i := 0; i < numDigits; i++ {
		output[i] = <-patternChan
	}
	doneChan <- true
	return output
}
