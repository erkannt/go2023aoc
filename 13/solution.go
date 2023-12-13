package main

import (
	"bufio"
	"log"
	"os"
)

func findReflection(input []string) int {
	for i := 0; i < len(input)-1; i++ {
		reflection := true
		for j := 0; i-j >= 0 && i+j+1 < len(input); j++ {
			reflection = (input[i+j+1] == input[i-j]) && reflection
		}
		if reflection {
			return i + 1
		}
	}
	return -1
}

func ProblemOne(scanner bufio.Scanner) int {
	total := 0
	orig := []string{}
	flipped := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			hLine := findReflection(orig)
			if hLine != -1 {
				total += 100 * hLine
			}

			vLine := findReflection(flipped)
			if vLine != -1 {
				total += vLine
			}

			orig = []string{}
			flipped = []string{}
			continue
		}

		orig = append(orig, line)
		if len(flipped) == 0 {
			for range line {
				flipped = append(flipped, "")
			}
		}
		for i, r := range line {
			flipped[i] += string(r)
		}
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
