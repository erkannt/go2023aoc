package main

import (
	"bufio"
	"strings"
	"testing"
)

type Location struct {
	x int
	y int
}

type Partnumber struct {
	value    int
	location Location
	lenght   int
}

func parseSchematic(scanner bufio.Scanner) ([]Partnumber, []Location) {
	return []Partnumber{}, []Location{}
}

func isAdjacent(number Partnumber, locations []Location) bool {
	return false
}

func ProblemOne(scanner bufio.Scanner) int {
	var total = 0
	partNumbers, symbolLocations := parseSchematic(scanner)
	for _, number := range partNumbers {
		if isAdjacent(number, symbolLocations) {
			total += number.value
		}
	}
	return total
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
