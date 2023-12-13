package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	actual := Solve(*bufio.NewScanner(strings.NewReader(input)), false)
	assert.Equal(t, expected, actual)
}
