package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkXPattern(data []string, rows, cols int) int {

	foundPatterns := 0
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if data[i][j] == 'A' {
				if checkXMASPattern(data, i, j) {
					foundPatterns++
				}
			}
		}
	}

	return foundPatterns
}

func checkXMASPattern(data []string, row, col int) bool {

	patternTuples := [][]string{{"MAS", "MAS"}, {"MAS", "SAM"}, {"SAM", "MAS"}, {"SAM", "SAM"}}

	for _, pattern := range patternTuples {
		ul := data[row-1][col-1]
		ur := data[row-1][col+1]

		ll := data[row+1][col-1]
		lr := data[row+1][col+1]

		if ul == pattern[0][0] && lr == pattern[0][2] && ur == pattern[1][0] && ll == pattern[1][2] {
			return true
		}
	}
	return false
}

func main() {

	data := make([]string, 0)

	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}

	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	rows, cols := len(data), len(data[0])
	count := checkXPattern(data, rows, cols)
	fmt.Println("Number of XMAS patterns found:", count)

}
