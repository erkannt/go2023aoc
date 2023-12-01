package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func onlyIfNumber(r rune) rune {
	if r <= '9' {
		return r
	}
	return -1
}

func onlyNumberRunes(input string) []rune {
	stringOfNumbers := strings.Map(onlyIfNumber, input)
	return []rune(stringOfNumbers)
}

func ProblemOne(input bufio.Scanner) int {
	total := 0
	for input.Scan() {
		line := input.Text()
		runes := onlyNumberRunes(line)
		if len(runes) == 0 {
			continue
		}

		firstAndLast := string([]rune{runes[0], runes[len(runes)-1]})
		value, err := strconv.Atoi(firstAndLast)
		if err != nil {
			println("Can't convert to number", line, firstAndLast)
			os.Exit(1)
		}

		total += value
	}
	return total
}

func ProblemTwo(input bufio.Scanner) int {
	return ProblemOne(input)
}

func main() {
	file, err := os.Open("./input/01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	println("Problem1:", ProblemOne(*bufio.NewScanner(file)))

	file.Seek(0, 0)
	println("Problem2:", ProblemTwo(*bufio.NewScanner(file)))
}
