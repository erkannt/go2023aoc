package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func PossibleArrangements(input string) int {
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
