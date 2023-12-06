package main

import (
	"bufio"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var problemInput = `
Time:      7  15   30
Distance:  9  40  200
`

type Race struct {
	duration int
	record   int
}

func parseRaces(scanner bufio.Scanner) []Race {
	numberRegex, _ := regexp.Compile("[0-9]+")
	races := []Race{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Time: ") {
			numbers := numberRegex.FindAllString(line, -1)
			for _, number := range numbers {
				value, _ := strconv.Atoi(number)
				races = append(races, Race{duration: value})
			}
		}
		if strings.HasPrefix(line, "Distance: ") {
			numbers := numberRegex.FindAllString(line, -1)
			for i, number := range numbers {
				value, _ := strconv.Atoi(number)
				races[i].record = value
			}
		}
	}
	return races
}

func optimiseForRace(race Race) (int, int) {
	durationSquared := math.Pow(float64(race.duration), 2)
	minButton := math.Floor(float64(race.duration)/2 - math.Sqrt(durationSquared/4-float64(race.record)))
	maxButton := math.Ceil(float64(race.duration)/2+math.Sqrt(durationSquared/4-float64(race.record))) - 1
	return int(minButton), int(maxButton)
}

func ProblemOne(scanner bufio.Scanner) int {
	total := 1
	races := parseRaces(scanner)
	for _, race := range races {
		minButton, maxButton := optimiseForRace(race)
		fmt.Printf("%v %v %v\n", race, minButton, maxButton)
		total *= maxButton - minButton
	}
	return total
}

func TestProblemOne(t *testing.T) {
	expected := 288

	result := ProblemOne(*bufio.NewScanner(strings.NewReader(problemInput)))
	if result != expected {
		t.Errorf("\n Expected: %v\nResult: %v", expected, result)
	}
}
