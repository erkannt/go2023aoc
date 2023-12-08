package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

type Node map[rune]string

func parseInput(scanner bufio.Scanner) ([]rune, map[string]Node) {
	instructions := []rune{}
	network := make(map[string]Node)
	nodeRegex := regexp.MustCompile("([A-Z]{3}).*([A-Z]{3}), ([A-Z]{3})")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "=") {
			matches := nodeRegex.FindAllStringSubmatch(line, -1)
			node := matches[0][1]
			left := matches[0][2]
			right := matches[0][3]
			network[node] = Node{'L': left, 'R': right}
			continue
		}
		instructions = []rune(line)
	}
	return instructions, network
}

func ProblemOne(scanner bufio.Scanner) int {
	instructions, network := parseInput(scanner)
	step := 0
	current := "AAA"
	for {
		instr := instructions[step%len(instructions)]
		current = network[current][instr]
		if current == "ZZZ" {
			break
		}
		step++
	}
	return step + 1
}

func ProblemTwo(scanner bufio.Scanner) int {
	step := 0
	return step + 1
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	println("Problem1:", ProblemOne(*bufio.NewScanner(file)))
}
