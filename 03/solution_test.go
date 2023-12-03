package main

import (
	"bufio"
	"reflect"
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
		t.Error("Wrong answer: ", result)
	}
}

func TestParseSchematic(t *testing.T) {
	input := `
467..114..
...*......
`
	expectedNumbers := []Partnumber{
		{value: 467,
			location: Location{x: 0, y: 0},
			lenght:   3,
		},
		{value: 114,
			location: Location{x: 5, y: 0},
			lenght:   3,
		},
	}
	expectedSymbolLocations := []Location{{x: 3, y: 1}}

	reader := bufio.NewScanner(strings.NewReader(input))
	resultNumbers, resultSymbolLocations := parseSchematic(*reader)
	if !reflect.DeepEqual(resultNumbers, expectedNumbers) {
		t.Error("Wrong answer: ", resultNumbers)
	}
	if !reflect.DeepEqual(resultSymbolLocations, expectedSymbolLocations) {
		t.Error("Wrong answer: ", resultSymbolLocations)
	}
}
