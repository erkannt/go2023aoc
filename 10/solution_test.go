package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestProblemOne(t *testing.T) {
	input := `.....
.S-7.
.|.|.
.L-J.
.....
`
	expected := 4
	actual := ProblemOne(*bufio.NewScanner(strings.NewReader(input)))
	if actual != expected {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestProblemTwo(t *testing.T) {
	input := `.....
.S-7.
.|.|.
.L-J.
.....
`
	expected := 1
	actual := ProblemTwo(*bufio.NewScanner(strings.NewReader(input)))
	if actual != expected {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}
