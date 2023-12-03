package main

import (
	"bufio"
	"fmt"
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

func TestProblemOneFromReddit(t *testing.T) {
	input := `
12.......*..
+.........34
.......-12..
..78........
..*....60...
78..........
.......23...
....90*12...
............
2.2......12.
.*.........*
1.1.......56
`
	expected := 413

	reader := bufio.NewScanner(strings.NewReader(input))
	result := ProblemOne(*reader)
	if result != expected {
		t.Error("Wrong answer: ", result)
	}
}

func TestProblemOneRealPartialInput(t *testing.T) {
	input := `
........954......104.......52......70..............206.806........708..........................217...............................440........
.......@...................*.............................*.664..............677................@....459.........687.........................
..................378.....398........548..495..........983....*................*..282.................*...........$.248.....409.......165...
`
	expected := 5897

	reader := bufio.NewScanner(strings.NewReader(input))
	result := ProblemOne(*reader)
	if result != expected {
		t.Error("Wrong answer: ", result)
	}
}

func TestMoreCases(t *testing.T) {
	cases := []struct {
		Input    string
		Expected int
	}{
		{`
........
.24..4..
......*.
`, 4},
		{`
........
.24$-4..
......*.
`, 28},
		{`
11....11
..$..$..
11....11
`, 44},
		{`
$......$
.1....1.
.1....1.
$......$
`, 4},
		{`
$......$
.11..11.
.11..11.
$......$
`, 44},
		{`
$11
...
11$
...
`, 22},
		{`
$..
.11
.11
$..
..$
11.
11.
..$
`, 44},
		{`
11.$.
`, 0},
	}
	for _, testCase := range cases {
		t.Run(fmt.Sprintf("%s -> %v", testCase.Input, testCase.Expected), func(t *testing.T) {
			result := ProblemOne(*bufio.NewScanner(strings.NewReader(testCase.Input)))
			if !reflect.DeepEqual(result, testCase.Expected) {
				t.Errorf("actual: %v", result)
			}
		})

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
