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

func ToNumbers(line string) []int {
	numbersByName := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	numbers := []int{}
	word := ""
	windowMin := 0
	for pos, char := range line {
		if char <= '9' {
			numbers = append(numbers, int(char-'0'))
			windowMin = pos + 1
			continue
		}
		for i := windowMin; i <= pos; i++ {
			word = line[i : pos+1]
			value := numbersByName[word]
			if value != 0 {
				numbers = append(numbers, value)
				break
			}
		}
	}
	return numbers
}

func ProblemTwo(input bufio.Scanner) int {
	total := 0
	for input.Scan() {
		line := input.Text()
		numbers := ToNumbers(line)
		if len(numbers) == 0 {
			continue
		}
		total += numbers[0]*10 + numbers[len(numbers)-1]
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
