package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func parse(input string) (string, []int) {
	numberRegex := regexp.MustCompile("[0-9]+")
	parts := strings.Split(input, " ")
	groupsStrings := numberRegex.FindAllStringSubmatch(parts[1], -1)
	groups := []int{}
	for _, v := range groupsStrings {
		value, _ := strconv.Atoi(v[0])
		groups = append(groups, value)
	}
	return parts[0], groups
}

func getCandidates(record string, groupSizes []int) [][]int {
	candidates := [][]int{}
	for _, size := range groupSizes {
		regx, _ := regexp.Compile(fmt.Sprintf("([?.][#?]{%d}[?.])", size))
		padded := fmt.Sprintf(".%s.", record)
		thisGroupsCandidates := []int{}
		offset := 0
		for {
			pos := regx.FindStringIndex(padded[offset:])
			if pos == nil {
				break
			}
			thisGroupsCandidates = append(thisGroupsCandidates, pos[0]+offset)
			offset += pos[0] + size
		}
		if len(thisGroupsCandidates) == 0 {
			log.Fatalf("Can't find position for group: %v %v %v", record, groupSizes, size)
		}
		candidates = append(candidates, thisGroupsCandidates)
	}
	return candidates
}

func remainingGroupArrangements(candidatePositions [][]int, sizes []int, startingPoint int) int {
	if len(candidatePositions) == 1 {
		possibleCount := 0
		for _, pos := range candidatePositions[0] {
			if pos >= startingPoint {
				possibleCount++
			}
		}
		return possibleCount
	}
	possibleCount := 0
	for _, pos := range candidatePositions[0] {
		if pos >= startingPoint {
			possibleCount += remainingGroupArrangements(candidatePositions[1:], sizes[1:], sizes[0]+pos+1)
		}
	}
	return possibleCount
}

func PossibleArrangements(input string) int {
	record, groupSizes := parse(input)
	candidatePositions := getCandidates(record, groupSizes)
	total := remainingGroupArrangements(candidatePositions, groupSizes, 0)
	return total
}

func TestPossibleArrangements(t *testing.T) {
	cases := []struct {
		Input    string
		Expected int
	}{
		{"???.### 1,1,3", 1},
		{".??..??...?##. 1,1,3", 4},
		{"?#?#?#?#?#?#?#? 1,3,1,6", 1},
		{"????.#...#... 4,1,1", 1},
		{"????.######..#####. 1,6,5", 4},
		{"?###???????? 3,2,1", 10},
	}
	for _, testCase := range cases {
		t.Run(fmt.Sprintf("%s -> %v", testCase.Input, testCase.Expected), func(t *testing.T) {
			assert.Equal(t, testCase.Expected, PossibleArrangements(testCase.Input))
		})
	}
}
