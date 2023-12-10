package main

import (
	"bufio"
	"fmt"
	"log"
	"slices"
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
	for y, line := range mapLines {
		for x, r := range line {
			if r == 'S' {
				return Position{x: x, y: y}
			}
		}
	}
	log.Fatal("Not start found")
	return Position{}
}

func findFirstStep(mapLines [][]rune, startPos Position) Step {
	switch {
	case slices.Contains([]rune{'-', 'L', 'F'}, mapLines[startPos.y][startPos.x-1]):
		return Step{x: -1, y: 0}
	case slices.Contains([]rune{'-', 'J', '7'}, mapLines[startPos.y][startPos.x+1]):
		return Step{x: 1, y: 0}
	case slices.Contains([]rune{'|', 'J', 'L'}, mapLines[startPos.y+1][startPos.x]):
		return Step{x: 0, y: 1}
	case slices.Contains([]rune{'|', 'F', '7'}, mapLines[startPos.y-1][startPos.x]):
		return Step{x: 0, y: -1}
	}
	log.Fatal("No valid first step found")
	return Step{}
}

func getNextStep(prevStep Step, pipeshape rune) Step {
	log.Printf("%v %v\n", prevStep, string(pipeshape))
	switch pipeshape {
	case '|':
		return Step{x: 0, y: prevStep.y}
	case '-':
		return Step{x: prevStep.x, y: 0}
	case 'L':
		if prevStep.y == 1 {
			return Step{x: 1, y: 0}
		}
		return Step{x: 0, y: -1}
	case 'J':
		if prevStep.y == 1 {
			return Step{x: -1, y: 0}
		}
		return Step{x: 0, y: -1}
	case '7':
		if prevStep.y == -1 {
			return Step{x: -1, y: 0}
		}
		return Step{x: 0, y: 1}
	case 'F':
		if prevStep.y == -1 {
			return Step{x: 1, y: 0}
		}
		return Step{x: 0, y: 1}
	}
	log.Fatal("Invalid pipeshape")
	return Step{}
}

func getPipeShape(pos Position, maplines [][]rune) rune {
	return maplines[pos.y][pos.x]
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
		fmt.Printf("%v  ", currentPos)
		if currentPos.x == startPos.x && currentPos.y == startPos.y {
			break
		}
		step = getNextStep(step, getPipeShape(currentPos, mapLines))
		pipeLength++
	}
	return pipeLength/2 + 1
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
