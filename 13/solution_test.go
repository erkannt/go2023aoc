package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ProblemOne(scanner bufio.Scanner) int {
	orig := []string{}
	flipped := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println("orig")
			for _, v := range orig {
				fmt.Printf("%v\n", v)
			}
			fmt.Println("flipped")
			for _, v := range flipped {
				fmt.Printf("%v\n", v)
			}
			orig = []string{}
			flipped = []string{}
			continue
		}

		orig = append(orig, line)
		if len(flipped) == 0 {
			for range line {
				flipped = append(flipped, "")
			}
		}
		for i, r := range line {
			flipped[i] += string(r)
		}
	}
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
