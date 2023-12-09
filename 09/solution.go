package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func parseHistories(scanner bufio.Scanner) [][]int {
	numbersRegex := regexp.MustCompile("[-0-9]+")
	histories := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		asStrings := numbersRegex.FindAllString(line, -1)
		numbers := []int{}
		for _, str := range asStrings {
			value, _ := strconv.Atoi(str)
			numbers = append(numbers, value)
		}
		histories = append(histories, numbers)
	}
	return histories
}

func isAllZero(input []int) bool {
	for _, v := range input {
		if v != 0 {
			return false
		}
	}
	return true
}

func nextValue(history []int) int {
	data := [][]int{}
	data = append(data, history)
	deltaCount := 0
	for {
		currentDeltas := []int{}
		for i := 1; i < len(data[deltaCount]); i++ {
			currentDeltas = append(currentDeltas, data[deltaCount][i]-data[deltaCount][i-1])
		}
		data = append(data, currentDeltas)
		deltaCount++
		if isAllZero(currentDeltas) {
			break
		}
	}
	predictionDelta := 0
	for i := deltaCount - 1; i >= 0; i-- {
		predictionDelta = data[i][len(data[i])-1] + predictionDelta
	}
	return predictionDelta
}

func ProblemOne(scanner bufio.Scanner) int {
	histories := parseHistories(scanner)
	total := 0
	for _, history := range histories {
		total += nextValue(history)
	}
	return total
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	println("Problem1:", ProblemOne(*bufio.NewScanner(file)))
	// file.Seek(0, 0)
	// println("Problem2:", ProblemTwo(*bufio.NewScanner(file)))
}
