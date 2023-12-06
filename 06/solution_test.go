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

func TestProblemOne(t *testing.T) {
	expected := 288

	result := ProblemOne(*bufio.NewScanner(strings.NewReader(problemInput)))
	if result != expected {
		t.Errorf("\n Expected: %v\nResult: %v", expected, result)
	}
}
