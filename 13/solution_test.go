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

func TestProblemTwo(t *testing.T) {
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
	expected := 400
	actual := Solve(*bufio.NewScanner(strings.NewReader(input)), true)
	assert.Equal(t, expected, actual)
}

func TestProblemDebugInput(t *testing.T) {
	input := `.##.###
.##.###
.######
##...##
#..#.#.
.##....
.##....
#.#....
...###.
...###.
#.#.#..
.##....
.##....
#..#.#.
##...##

`
	expected := 900
	actual := Solve(*bufio.NewScanner(strings.NewReader(input)), true)
	assert.Equal(t, expected, actual)
}
