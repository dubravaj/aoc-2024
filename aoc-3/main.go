package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
)

func main() {

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	content := string(data)

	// part 1

	var mulRe = regexp.MustCompile(`(mul)\((\d{1,3}),(\d{1,3})\)`)
	matches := mulRe.FindAllString(content, -1)

	sum := 0
	var a, b int
	for _, match := range matches {
		fmt.Sscanf(match, "mul(%d,%d)", &a, &b)
		sum += a * b
	}

	fmt.Println("Sum of all mul operations:", sum)

	// part 2
	doOpRegex := regexp.MustCompile(`(do)\(\)`)
	dontOpRegex := regexp.MustCompile(`(don\'t)\(\)`)

	opPatterns := []*regexp.Regexp{mulRe, doOpRegex, dontOpRegex}

	opMatchesPositions := make([][]int, 0)

	for _, opPattern := range opPatterns {
		matches := opPattern.FindAllStringIndex(content, -1)
		opMatchesPositions = append(opMatchesPositions, matches...)
	}

	sort.Slice(opMatchesPositions, func(i, j int) bool {
		return opMatchesPositions[i][0] < opMatchesPositions[j][0]
	})

	doMultiplication := true
	sum = 0

	for _, matchPosition := range opMatchesPositions {
		matchOp := content[matchPosition[0]:matchPosition[1]]
		op := matchOp[:3]

		switch op {
		case "mul":
			if doMultiplication {
				fmt.Sscanf(matchOp, "mul(%d,%d)", &a, &b)
				sum += a * b
			}
		case "don":
			doMultiplication = false
		default:
			doMultiplication = true
		}
	}

	fmt.Println("Sum of all mul operations with do() and don't() operations:", sum)

}
