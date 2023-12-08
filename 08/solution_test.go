package main

import (
	"bufio"
	"strings"
	"testing"
)

var problemInput = `
LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`

func TestProblemOne(t *testing.T) {
	expected := 6

	result := ProblemOne(*bufio.NewScanner(strings.NewReader(problemInput)))
	if result != expected {
		t.Errorf("\n Expected: %v\nResult: %v", expected, result)
	}
}
