package main

import (
	"fmt"
	"testing"
)

func ProblemOne(input string) int {
	fmt.Println("Hello, world.")
	return 42
}

func TestProblemOne(t *testing.T) {
	input := `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`
	ans := ProblemOne(input)
	if ans != 142 {
		t.Error("wrong result", ans)
	}
}
