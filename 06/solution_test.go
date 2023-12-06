package main

import (
	"bufio"
	"strings"
	"testing"
)

var problemInput = `
Time:      7  15   30
Distance:  9  40  200
`

type Race struct {
	duration int
	record   int
}

func parseRaces(scanner bufio.Scanner) []Race {
	return []Race{}
}

func optimiseForRace(race Race) (int, int) {
	return 0, 0
}

func ProblemOne(scanner bufio.Scanner) int {
	total := 0
	races := parseRaces(scanner)
	for _, race := range races {
		minButton, maxButton := optimiseForRace(race)
		total *= maxButton - minButton
	}
	return total
}

func TestProblemOne(t *testing.T) {
	expected := 35

	result := ProblemOne(*bufio.NewScanner(strings.NewReader(problemInput)))
	if result != expected {
		t.Errorf("\n Expected: %v\nResult: %v", expected, result)
	}
}
