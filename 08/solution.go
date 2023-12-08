package main

import (
	"bufio"
	"log"
	"os"
)

func ProblemOne(scanner bufio.Scanner) int {
	total := 1
	return total
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	println("Problem1:", ProblemOne(*bufio.NewScanner(file)))
}
