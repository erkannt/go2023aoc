package main

import (
	"testing"
)

func TestProblemOne(t *testing.T) {
	input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	expected := 374
	actual := ProblemOne(input)
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\v", expected, actual)
	}
}

func TestProblemTwo(t *testing.T) {
	input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	expected := 1030
	actual := ProblemTwo(input, 10)
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\v", expected, actual)
	}
}
