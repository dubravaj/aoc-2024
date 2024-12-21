package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func IsValid(sequence []int) bool {

	if len(sequence) == 0 {
		return false
	}

	// set decreasing and increasing according to the first two elements
	var increasing bool

	if sequence[0] < sequence[1] {
		increasing = true
	} else {
		increasing = false
	}

	for i := 1; i < len(sequence); i++ {
		lastLevel := sequence[i-1]
		currentLevel := sequence[i]

		levelDiff := currentLevel - lastLevel

		if math.Abs(float64(levelDiff)) < 1.0 || math.Abs(float64(levelDiff)) > 3.0 {
			return false
		}

		// sequence changed direction
		if (increasing && lastLevel > currentLevel) || (!increasing && lastLevel < currentLevel) {
			return false
		}

	}
	return true
}

func CheckOneLevelError(sequence []int, index int) bool {
	if index >= len(sequence) {
		return false
	}

	sequenceCopy := make([]int, len(sequence))
	copy(sequenceCopy, sequence)

	sequenceRemoved := append(sequenceCopy[:index], sequenceCopy[index+1:]...)

	if IsValid(sequenceRemoved) {
		return true
	} else {
		return CheckOneLevelError(sequence, index+1)
	}
}

func main() {

	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}

	scanner := bufio.NewScanner(inputFile)

	levelsSequence := make([][]int, 0)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		sequence := make([]int, 0)
		for _, num := range line {
			numInt, _ := strconv.Atoi(num)
			sequence = append(sequence, numInt)
		}

		levelsSequence = append(levelsSequence, sequence)

	}

	safeSequences := 0

	for _, sequence := range levelsSequence {
		if IsValid(sequence) {
			safeSequences++
		} else {
			// try to continue with invalid sequence and remove one element
			hasOneLevelError := CheckOneLevelError(sequence, 0)
			if hasOneLevelError {
				safeSequences++
			}
		}
	}

	fmt.Println("Number of safe sequence: ", safeSequences)
}
