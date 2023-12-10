package main

import (
	"testing"
)

func ProblemOne(input string) int {
	return 0
}

func TestProblemOne(t *testing.T) {
	input := `.....
.S-7.
.|.|.
.L-J.
.....
`
	expected := 4
	actual := ProblemOne(input)
	if actual != expected {
		t.Errorf("Expected: %v\nActual: %v\n", expected, actual)
	}
}
