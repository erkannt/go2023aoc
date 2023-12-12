package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		{"....#?##????.??#?? 4,1", 1},
	}
	for _, testCase := range cases {
		t.Run(fmt.Sprintf("%s -> %v", testCase.Input, testCase.Expected), func(t *testing.T) {
			assert.Equal(t, testCase.Expected, PossibleArrangements(testCase.Input))
		})
	}
}
