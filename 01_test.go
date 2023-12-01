package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func onlyIfNumber(r rune) rune {
	if r <= '9' {
		return r
	}
	return -1
}

func ProblemOne(input bufio.Scanner) int {
	for input.Scan() {
		line := input.Text()
		numbers := strings.Map(onlyIfNumber, line)
		if numbers == "" {
			continue
		}
		value, err := strconv.Atoi(numbers)
		if err != nil {
			println("Can't convert to number", numbers)
			os.Exit(1)
		}
		return value
	}
	println("no values found")
	return -1
}

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
