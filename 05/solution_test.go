package main

import (
	"bufio"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

type Seed int

type SeedMapping struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}

func parseAlmanac(scanner bufio.Scanner) ([]Seed, [][]SeedMapping) {
	numberRegex, _ := regexp.Compile("[0-9]+")
	seeds := []Seed{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.Contains(line, "seeds: ") {
			seedStrings := numberRegex.FindAllString(line, -1)
			for _, str := range seedStrings {
				value, _ := strconv.Atoi(str)
				seeds = append(seeds, Seed(value))
			}

		}
	}
	return seeds, [][]SeedMapping{}
}

func lookupLocation(seed Seed, maps [][]SeedMapping) int {
	return 0
}

func ProblemOne(scanner bufio.Scanner) int {
	seeds, maps := parseAlmanac(scanner)
	var nearestLocation = lookupLocation(seeds[0], maps)
	for _, seed := range seeds {
		location := lookupLocation(seed, maps)
		if location < nearestLocation {
			nearestLocation = location
		}
	}
	return nearestLocation
}
func TestParseAlmanacSeeds(t *testing.T) {
	input := `
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15
`
	expected := []Seed{79, 14, 55, 13}

	result, _ := parseAlmanac(*bufio.NewScanner(strings.NewReader(input)))
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\n Expected: %v\nResult: %v", expected, result)
	}
}

func TestProblemOne(t *testing.T) {
	input := `
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`
	expected := 35

	result := ProblemOne(*bufio.NewScanner(strings.NewReader(input)))
	if result != expected {
		t.Errorf("\n Expected: %v\nResult: %v", expected, result)
	}
}
