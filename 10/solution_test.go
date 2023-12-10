package main

import (
	"bufio"
	"strings"
	"testing"
)

type Position struct {
	x int
	y int
}

type Step struct {
	x int
	y int
}

func findStart(mapLines [][]rune) Position {
	return Position{}
}

func findFirstStep(mapLines [][]rune, startPos Position) Step {
	return Step{}
}

func getNextStep(prevStep Step, pipeshape rune) Step {
	return Step{}
}

func getPipeShape(pos Position, maplines [][]rune) rune {
	return '|'
}

func ProblemOne(scanner bufio.Scanner) int {
	mapLines := [][]rune{}
	for scanner.Scan() {
		mapLines = append(mapLines, []rune(scanner.Text()))
	}
	pipeLength := 0
	startPos := findStart(mapLines)
	currentPos := findStart(mapLines)
	step := findFirstStep(mapLines, currentPos)
	for {
		currentPos.x += step.x
		currentPos.y += step.y
		if currentPos.x == startPos.x && currentPos.y == startPos.y {
			break
		}
		step = getNextStep(step, getPipeShape(currentPos, mapLines))
		pipeLength++
	}
	return pipeLength
}

func TestProblemOne(t *testing.T) {
	input := `.....
.S-7.
.|.|.
.L-J.
.....
`
	expected := 4
	actual := ProblemOne(*bufio.NewScanner(strings.NewReader(input)))
	if actual != expected {
		t.Errorf("Expected: %v\nActual: %v\n", expected, actual)
	}
}
