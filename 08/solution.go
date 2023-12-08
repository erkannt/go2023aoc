package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Node map[rune]string

func parseInput(scanner bufio.Scanner) ([]rune, map[string]Node) {
	instructions := []rune{}
	network := make(map[string]Node)
	nodeRegex := regexp.MustCompile("([A-Z0-9]{3}).*([A-Z0-9]{3}), ([A-Z0-9]{3})")
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

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func ProblemTwo(scanner bufio.Scanner) int {
	instructions, network := parseInput(scanner)
	current := []string{}
	for k := range network {
		if strings.HasSuffix(k, "A") {
			current = append(current, k)
		}
	}
	fmt.Printf("%v\n", current)

	steps := []int{}

	for _, node := range current {
		step := 0
		for {
			instr := instructions[step%len(instructions)]
			node = network[node][instr]
			if strings.HasSuffix(node, "Z") {
				steps = append(steps, step+1)
				break
			}
			step++
		}
	}
	fmt.Printf("%v\n", steps)
	return LCM(steps[0], steps[1], steps...)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	println("Problem1:", ProblemOne(*bufio.NewScanner(file)))
	file.Seek(0, 0)
	println("Problem2:", ProblemTwo(*bufio.NewScanner(file)))
}
