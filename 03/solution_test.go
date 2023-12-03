package main

import (
	"bufio"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
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
	symbolRegx, _ := regexp.Compile("[^0-9.]")
	numberRegx, _ := regexp.Compile("[0-9]+")

	var numbers = []Partnumber{}
	var locations = []Location{}

	var y = 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		symbolLocs := symbolRegx.FindAllIndex([]byte(line), -1)
		for _, loc := range symbolLocs {
			locations = append(locations, Location{x: loc[0], y: y})
		}

		numberLocs := numberRegx.FindAllIndex([]byte(line), -1)
		for _, loc := range numberLocs {
			value, _ := strconv.Atoi(line[loc[0]:loc[1]])
			newNumber := Partnumber{
				value:  value,
				lenght: loc[1] - loc[0],
				location: Location{
					x: loc[0],
					y: y,
				},
			}
			numbers = append(numbers, newNumber)
		}
		y += 1
	}

	return numbers, locations
}

func isAdjacent(number Partnumber, locations []Location) bool {
	var neighbours = []string{
		fmt.Sprintf("%d%d", number.location.x-1, number.location.y),
		fmt.Sprintf("%d%d", number.location.x+1, number.location.y),
	}
	for i := number.location.x - 1; i <= number.location.x+number.lenght+1; i++ {
		neighbours = append(neighbours, fmt.Sprintf("%d%d", i, number.location.y-1))
		neighbours = append(neighbours, fmt.Sprintf("%d%d", i, number.location.y+1))
	}
	for _, loc := range locations {
		for _, neighbour := range neighbours {
			if neighbour == fmt.Sprintf("%d%d", loc.x, loc.y) {
				return true
			}
		}
	}
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
