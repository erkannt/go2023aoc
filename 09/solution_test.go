package main

import (
	"bufio"
	"strings"
	"testing"
)

func ProblemOne(scanner bufio.Scanner) int {
	return 0
}

func TestProblemOne(t *testing.T) {
	input := `
0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`

	expected := 114

	result := ProblemOne(*bufio.NewScanner(strings.NewReader(input)))
	if result != expected {
		t.Errorf("\n Expected: %v\nResult: %v", expected, result)
	}
}
