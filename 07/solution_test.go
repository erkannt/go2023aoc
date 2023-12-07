package main

import (
	"bufio"
	"strings"
	"testing"
)

var problemInput = `
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

var problemInputReddit = `
2345A 1
Q2KJJ 13
Q2Q2Q 19
T3T3J 17
T3Q33 11
2345J 3
J345A 2
32T3K 5
T55J5 29
KK677 7
KTJJT 34
QQQJA 31
JJJJJ 37
JAAAA 43
AAAAJ 59
AAAAA 61
2AAAA 23
2JJJJ 53
JJJJ2 41
`

func TestProblemOne(t *testing.T) {
	expected := 6440

	result := ProblemOne(*bufio.NewScanner(strings.NewReader(problemInput)))
	if result != expected {
		t.Errorf("\n Expected: %v\nResult: %v", expected, result)
	}
}

func TestProblemOneReddit(t *testing.T) {
	expected := 6592

	result := ProblemOne(*bufio.NewScanner(strings.NewReader(problemInputReddit)))
	if result != expected {
		t.Errorf("\n Expected: %v\nResult: %v", expected, result)
	}
}
