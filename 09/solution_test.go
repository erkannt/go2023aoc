package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func parseHistories(scanner bufio.Scanner) [][]int {
	numbersRegex := regexp.MustCompile("[-0-9]+")
	histories := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		asStrings := numbersRegex.FindAllString(line, -1)
		numbers := []int{}
		for _, str := range asStrings {
			value, _ := strconv.Atoi(str)
			numbers = append(numbers, value)
		}
		fmt.Printf("%v\n", numbers)
		histories = append(histories, numbers)
	}
	return histories
}

func ProblemOne(scanner bufio.Scanner) int {
	histories := parseHistories(scanner)
	fmt.Printf("%v\n", histories)
	return 0
}

func TestProblemOne(t *testing.T) {
	input := `
0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`

	expected := 114

	result := ProblemOne(*bufio.NewScanner(strings.NewReader(input)))
	if result != expected {
		t.Errorf("\n Expected: %v\nResult: %v", expected, result)
	}
}
