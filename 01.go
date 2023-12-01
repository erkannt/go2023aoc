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

func toNumbers(line string) []int {
	numbers := []int{}
	for _, char := range line {
		if char <= '9' {
			numbers = append(numbers, int(char-'0'))
			continue
		}
	}
	return numbers
}

func ProblemTwo(input bufio.Scanner) int {
	total := 0
	for input.Scan() {
		line := input.Text()
		numbers := toNumbers(line)
		switch len(numbers) {
		case 0:
			continue
		case 1:
			total += numbers[0]
		default:
			total += numbers[0]*10 + numbers[len(numbers)-1]
		}
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
