package main

import (
	"bufio"
	"strings"
	"testing"
)

var problemInput = `
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

func ProblemOne(scanner bufio.Scanner) int {
	return 0
}

func TestProblemOne(t *testing.T) {
	expected := 6440

	result := ProblemOne(*bufio.NewScanner(strings.NewReader(problemInput)))
	if result != expected {
		t.Errorf("\n Expected: %v\nResult: %v", expected, result)
	}
}
