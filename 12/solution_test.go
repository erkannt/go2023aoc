package main

import (
	"fmt"
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

func PossibleArrangements(input string) int {
	record, groups := parse(input)
	fmt.Printf("%v %v", record, groups)
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
