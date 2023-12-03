package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

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

func TestIsAdjacent1(t *testing.T) {
	result := isAdjacent(
		Partnumber{
			value:    467,
			location: Location{x: 0, y: 0},
			lenght:   3,
		},
		[]Location{{x: 1, y: 1}},
	)
	if result != true {
		t.Error("Wrong answer")
	}
}

func TestIsAdjacent2(t *testing.T) {
	result := isAdjacent(
		Partnumber{value: 114,
			location: Location{x: 5, y: 0},
			lenght:   3,
		},
		[]Location{{x: 1, y: 1}},
	)
	if result != false {
		t.Error("Wrong answer")
	}
}
