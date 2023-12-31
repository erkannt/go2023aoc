package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

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

func TestParseAlmanacMaps(t *testing.T) {
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
	expected := [][]SeedMapping{
		{
			SeedMapping{
				destinationStart: 50,
				sourceStart:      98,
				rangeLength:      2,
			},
			SeedMapping{
				destinationStart: 52,
				sourceStart:      50,
				rangeLength:      48,
			},
		},
		{
			SeedMapping{
				destinationStart: 0,
				sourceStart:      15,
				rangeLength:      37,
			},
			SeedMapping{
				destinationStart: 37,
				sourceStart:      52,
				rangeLength:      2,
			},
			SeedMapping{
				destinationStart: 39,
				sourceStart:      0,
				rangeLength:      15,
			},
		},
	}

	_, result := parseAlmanac(*bufio.NewScanner(strings.NewReader(input)))
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\n Expected: %v\nResult: %v", expected, result)
	}
}

var problemInput = `
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

func TestProblemOne(t *testing.T) {
	expected := 35

	result := ProblemOne(*bufio.NewScanner(strings.NewReader(problemInput)))
	if result != expected {
		t.Errorf("\n Expected: %v\nResult: %v", expected, result)
	}
}
func TestProblemTwo(t *testing.T) {
	expected := 46

	result := ProblemTwo(*bufio.NewScanner(strings.NewReader(problemInput)))
	if result != expected {
		t.Errorf("\n Expected: %v\nResult: %v", expected, result)
	}
}
