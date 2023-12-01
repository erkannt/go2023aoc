package main

import (
	"bufio"
	"strings"
	"testing"
)

func ProblemOne(input bufio.Reader) int {
	return 42
}

func TestProblemOne(t *testing.T) {
	input := `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

	reader := bufio.NewReader(strings.NewReader(input))
	ans := ProblemOne(*reader)
	if ans != 142 {
		t.Error("wrong result", ans)
	}
}
