package main

import (
	"bufio"
	"strings"
	"testing"
)

func ProblemOne(bufio.Scanner) int {
	return 0
}

func TestProblemOne(t *testing.T) {
	input := `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`
	expected := 4361

	reader := bufio.NewScanner(strings.NewReader(input))
	result := ProblemOne(*reader)
	if result != expected {
		t.Fatal("Wrong answer: ", result)
	}
}
