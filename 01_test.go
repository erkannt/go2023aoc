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
	total := 0
	for input.Scan() {
		line := input.Text()
		numbers := strings.Map(onlyIfNumber, line)
		if numbers == "" {
			continue
		}
		runes := []rune(numbers)
		firstAndLast := string([]rune{runes[0], runes[len(runes)-1]})
		println(firstAndLast)
		value, err := strconv.Atoi(firstAndLast)
		if err != nil {
			println("Can't convert to number", numbers)
			os.Exit(1)
		}
		total += value
	}
	println("no values found")
	return total
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
