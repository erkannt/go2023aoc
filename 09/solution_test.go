package main

import (
	"bufio"
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
		histories = append(histories, numbers)
	}
	return histories
}

func sum(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total
}

func nextValue(history []int) int {
	data := [][]int{}
	data = append(data, history)
	deltaCount := 0
	for {
		currentDeltas := []int{}
		for i := 1; i < len(data[deltaCount]); i++ {
			currentDeltas = append(currentDeltas, data[deltaCount][i]-data[deltaCount][i-1])
		}
		data = append(data, currentDeltas)
		deltaCount++
		if sum(currentDeltas) == 0 {
			break
		}
	}
	predictionDelta := 0
	for i := deltaCount - 1; i >= 0; i-- {
		predictionDelta = data[i][len(data[i])-1] + predictionDelta
	}
	return predictionDelta
}

func ProblemOne(scanner bufio.Scanner) int {
	histories := parseHistories(scanner)
	total := 0
	for _, history := range histories {
		total += nextValue(history)
	}
	return total
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
