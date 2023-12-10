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

func TestProblemTwoComplexExample(t *testing.T) {
	input := `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...
`

	expected := 8
	actual := ProblemTwo(*bufio.NewScanner(strings.NewReader(input)))
	if actual != expected {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}
