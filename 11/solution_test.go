package main

import "testing"

func ProblemOne(input string) int {
	return 0
}

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
