package main

import (
	"bufio"
	"fmt"
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

func toNumbers(line string) []int {
	return []int{}
}

func ProblemTwo(input bufio.Scanner) int {
	total := 0
	for input.Scan() {
		line := input.Text()
		numbers := toNumbers(line)
		if len(numbers) == 0 {
			continue
		}
		firstAndLast := fmt.Sprintf("%d%d", numbers[0], numbers[len(numbers)-1])
		value, err := strconv.Atoi(firstAndLast)
		if err != nil {
			println("Can't convert to number", line, firstAndLast)
			os.Exit(1)
		}
		total += value
	}
	return total
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
