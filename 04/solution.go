package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
)

func parseCard(card string) ([]string, []string) {
	sectionsRegex, _ := regexp.Compile("([0-9]+):([0-9 ]+)|([0-9 ]+)")
	numberRegex, _ := regexp.Compile("[0-9]+")

	sections := sectionsRegex.FindAllStringSubmatch(card, -1)
	winningNumbers := numberRegex.FindAllString(sections[1][0], -1)
	cardNumbers := numberRegex.FindAllString(sections[2][0], -1)
	return winningNumbers, cardNumbers
}

func cardValue(matchCount int) int {
	if matchCount == 0 {
		return 0
	}
	value := math.Pow(2, float64(matchCount)-1)
	return int(value)
}

func ProblemOne(scanner bufio.Scanner) int {
	var total = 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		winningNumbers, cardNumbers := parseCard(line)
		var matches = []string{}
		for _, cardNumber := range cardNumbers {
			for _, winningNumber := range winningNumbers {
				if cardNumber == winningNumber {
					matches = append(matches, cardNumber)
				}
			}
		}
		total += cardValue(len(matches))
	}
	return total
}

func ProblemTwo(scanner bufio.Scanner) int {
	return 0
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	println("Problem1:", ProblemOne(*bufio.NewScanner(file)))
}
