package main

import (
	"bufio"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func parseCard(card string) ([]string, []string) {
	sectionsRegex, _ := regexp.Compile("(Card [0-9]+: )([0-9 ]+)|([0-9]+)")
	numberRegex, _ := regexp.Compile("[0-9]")

	sections := sectionsRegex.FindAllStringSubmatch(card, -1)
	winningNumbers := numberRegex.FindAllString(sections[1][0], -1)
	cardNumbers := numberRegex.FindAllString(sections[2][0], -1)
	return winningNumbers, cardNumbers
}

func cardValue(matchCount int) int {
	if matchCount == 0 {
		return 0
	}
	value := math.Pow(2, float64(matchCount)-1)
	return int(value)
}

func ProblemOne(scanner bufio.Scanner) int {
	var total = 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		winningNumbers, cardNumbers := parseCard(line)
		var matches = []string{}
		for _, cardNumber := range cardNumbers {
			for _, winningNumber := range winningNumbers {
				if cardNumber == winningNumber {
					matches = append(matches, cardNumber)
				}
			}
		}
		total += cardValue(len(matches))
	}
	return total
}

func TestProblemOne(t *testing.T) {
	input := `
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`
	expected := 13
	result := ProblemOne(*bufio.NewScanner(strings.NewReader(input)))

	if expected != result {
		t.Errorf("\nExpected: %v\nResult: %v", expected, result)
	}

}

func TestCardValue(t *testing.T) {
	cases := []struct {
		Input    int
		Expected int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 8},
	}
	for _, testCase := range cases {
		t.Run(fmt.Sprintf("%v -> %v", testCase.Input, testCase.Expected), func(t *testing.T) {
			result := cardValue(testCase.Input)
			if !reflect.DeepEqual(result, testCase.Expected) {
				t.Errorf("actual: %v", result)
			}
		})

	}
}
