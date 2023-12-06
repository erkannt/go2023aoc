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

func ProblemOne(scanner bufio.Scanner) int {
	return 0
}

func TestProblemOne(t *testing.T) {
	expected := 35

	result := ProblemOne(*bufio.NewScanner(strings.NewReader(problemInput)))
	if result != expected {
		t.Errorf("\n Expected: %v\nResult: %v", expected, result)
	}
}
