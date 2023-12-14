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

func findUnsmudgedValue(orig []string, flipped []string) int {
	hLine := findReflection(orig)
	if hLine != -1 {
		return 100 * hLine
	}

	vLine := findReflection(flipped)
	return vLine
}

func findSmudgedValue(orig []string, flipped []string, unsmudged int) int {
	asRunes := [][]rune{}
	for _, line := range orig {
		asRunes = append(asRunes, []rune(line))
	}

	for i := 0; i < len(orig); i++ {
		for j := 0; j < len(orig[0]); j++ {
			smudged := orig
			smudgedLine := asRunes[i]
			if smudged[i][j] == '#' {
				smudgedLine[j] = '.'
			} else {
				smudgedLine[j] = '#'
			}
			smudged[i] = string(smudgedLine)
			hLine := findReflection(smudged)
			if hLine != -1 && hLine*100 != unsmudged {
				return 100 * hLine
			}
		}
	}

	asRunes = [][]rune{}
	for _, line := range flipped {
		asRunes = append(asRunes, []rune(line))
	}

	for i := 0; i < len(flipped); i++ {
		for j := 0; j < len(flipped[0]); j++ {
			smudged := flipped
			smudgedLine := asRunes[i]
			if smudged[i][j] == '#' {
				smudgedLine[j] = '.'
			} else {
				smudgedLine[j] = '#'
			}
			smudged[i] = string(smudgedLine)
			vLine := findReflection(smudged)
			if vLine != -1 && vLine != unsmudged {
				return vLine
			}
		}
	}
	return -1
}

func Solve(scanner bufio.Scanner, smudge bool) int {
	total := 0
	orig := []string{}
	flipped := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			unsmudged := findUnsmudgedValue(orig, flipped)
			if smudge {
				total += findSmudgedValue(orig, flipped, unsmudged)
			} else {
				total += unsmudged
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

	println("Problem1:", Solve(*bufio.NewScanner(file), false))
	file.Seek(0, 0)
	println("Problem2:", Solve(*bufio.NewScanner(file), true))
}
