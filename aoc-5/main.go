package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type PageRule struct {
	Page       int
	BeforePage int
}

func parseInput(input []string) ([]PageRule, [][]int) {
	pageRules := make([]PageRule, 0)
	pageSequences := make([][]int, 0)
	pageRuleRe := regexp.MustCompile(`(\d{1,2})\|(\d{1,2})`)

	for _, line := range input {
		// skipt empty line separating page rules and page sequences
		if line == "" {
			continue
		}
		if pageRuleRe.MatchString(line) {
			pageRule := PageRule{}
			fmt.Sscanf(line, "%d|%d", &pageRule.Page, &pageRule.BeforePage)
			pageRules = append(pageRules, pageRule)
		} else {
			sequence := make([]int, 0)
			// split line by comma
			sequenceStr := strings.Split(line, ",")
			for _, page := range sequenceStr {
				pageInt, _ := strconv.Atoi(page)
				sequence = append(sequence, pageInt)
			}
			pageSequences = append(pageSequences, sequence)
		}
	}
	return pageRules, pageSequences
}

func createPageOrdering(pageRules []PageRule) map[int][]int {
	pageOrderingMap := make(map[int][]int)

	for _, pageRule := range pageRules {
		pageOrderingMap[pageRule.Page] = append(pageOrderingMap[pageRule.Page], pageRule.BeforePage)
	}

	for _, pageOrdering := range pageOrderingMap {
		sort.Ints(pageOrdering)
	}

	return pageOrderingMap
}

func validatePageOrdering(pageOrdering map[int][]int, sequence []int) bool {

	for i := 0; i < len(sequence)-1; i++ {
		page := sequence[i]
		pageBefore := sequence[i+1]

		index := sort.SearchInts(pageOrdering[page], pageBefore)
		if index == len(pageOrdering[page]) || pageOrdering[page][index] != pageBefore {
			return false
		}
	}

	return true
}

func invalidSequenceToValid(pageOrdering map[int][]int, sequence []int) []int {
	validSequence := make([]int, len(sequence))
	copy(validSequence, sequence)

	for !validatePageOrdering(pageOrdering, validSequence) {
		for i := 0; i < len(validSequence)-1; i++ {
			page := validSequence[i]
			pageBefore := validSequence[i+1]

			index := sort.SearchInts(pageOrdering[page], pageBefore)
			if index == len(pageOrdering[page]) || pageOrdering[page][index] != pageBefore {
				// swap pages
				validSequence[i], validSequence[i+1] = validSequence[i+1], validSequence[i]
				break
			}
		}
	}

	return validSequence
}

func main() {

	data := make([]string, 0)

	inputFile, err := os.Open("input")
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

	pageRules, pageSequences := parseInput(data)
	pageOrdering := createPageOrdering(pageRules)

	validSequences := make([][]int, 0)
	invalidSequences := make([][]int, 0)
	validSum := 0
	validFromInvalidSum := 0

	for _, sequence := range pageSequences {
		if validatePageOrdering(pageOrdering, sequence) {
			validSequences = append(validSequences, sequence)
		} else {
			invalidSequences = append(invalidSequences, sequence)
		}
	}

	// part 1
	for _, sequence := range validSequences {
		middlePageIndex := len(sequence) / 2
		middlePage := sequence[middlePageIndex]
		validSum += middlePage
	}

	// part 2
	for _, sequence := range invalidSequences {
		validSequence := invalidSequenceToValid(pageOrdering, sequence)
		middlePageIndex := len(validSequence) / 2
		validFromInvalidSum += validSequence[middlePageIndex]
	}

	fmt.Println("Part 1: ", validSum)
	fmt.Println("Part 2: ", validFromInvalidSum)
}
