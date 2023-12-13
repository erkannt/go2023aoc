package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ProblemOne(scanner bufio.Scanner) int {
	return 0
}

func TestProblemOne(t *testing.T) {
	input := `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#

`
	expected := 405
	actual := ProblemOne(*bufio.NewScanner(strings.NewReader(input)))
	assert.Equal(t, expected, actual)
}
