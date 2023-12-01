package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestProblemOne(t *testing.T) {
	input := `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

	reader := bufio.NewScanner(strings.NewReader(input))
	ans := ProblemOne(*reader)
	if ans != 142 {
		t.Error("wrong result", ans)
	}
}

func TestProblemTwo(t *testing.T) {
	input := `
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

	reader := bufio.NewScanner(strings.NewReader(input))
	ans := ProblemTwo(*reader)
	if ans != 281 {
		t.Error("wrong result", ans)
	}
}

func TestToNumbers(t *testing.T) {
	input := "zoneight234"
	expected := []int{1, 8, 2, 3, 4}
	result := ToNumbers(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Input: %v\nResult: %v\nExpected %v", input, result, expected)
	}
}
