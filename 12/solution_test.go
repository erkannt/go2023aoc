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
		regx, _ := regexp.Compile(fmt.Sprintf("[?.][#?]{%d}[?.]", size))
		padded := fmt.Sprintf(".%s.", record)
		thisGroupsCandidates := regx.FindAllSubmatchIndex([]byte(padded), -1)
		if thisGroupsCandidates == nil {
			log.Fatalf("Can't find position for group: %v %v %v", record, groupSizes, size)
		}
		candidates = append(candidates, thisGroupsCandidates...)
	}
	return candidates
}

func PossibleArrangements(input string) int {
	record, groupSizes := parse(input)
	candidatePositions := getCandidates(record, groupSizes)
	for _, v := range candidatePositions {
		fmt.Printf("%v\n", v)
	}
	return 0
}

func TestPossibleArrangements(t *testing.T) {
	cases := []struct {
		Input    string
		Expected int
	}{
		{"???.### 1,1,3", 1},
	}
	for _, testCase := range cases {
		t.Run(fmt.Sprintf("%s -> %v", testCase.Input, testCase.Expected), func(t *testing.T) {
			assert.Equal(t, PossibleArrangements(testCase.Input), testCase.Expected)
		})
	}
}
