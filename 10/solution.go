package main

import (
	"bufio"
	"fmt"
	"log"
	"maps"
	"os"
	"slices"
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

type Turn = int

const (
	NoTurn           Turn = 0
	Clockwise        Turn = 1
	Counterclockwise Turn = -1
)

func getNextStep(prevStep Step, pipeshape rune) (Step, Turn) {
	switch pipeshape {
	case '|':
		return Step{x: 0, y: prevStep.y}, NoTurn
	case '-':
		return Step{x: prevStep.x, y: 0}, NoTurn
	case 'L':
		if prevStep.y == 1 {
			return Step{x: 1, y: 0}, Counterclockwise
		}
		return Step{x: 0, y: -1}, Clockwise
	case 'J':
		if prevStep.y == 1 {
			return Step{x: -1, y: 0}, Clockwise
		}
		return Step{x: 0, y: -1}, Counterclockwise
	case '7':
		if prevStep.y == -1 {
			return Step{x: -1, y: 0}, Counterclockwise
		}
		return Step{x: 0, y: 1}, Clockwise
	case 'F':
		if prevStep.y == -1 {
			return Step{x: 1, y: 0}, Clockwise
		}
		return Step{x: 0, y: 1}, Counterclockwise
	}
	log.Fatal("Invalid pipeshape")
	return Step{}, NoTurn
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
		if currentPos.x == startPos.x && currentPos.y == startPos.y {
			break
		}
		step, _ = getNextStep(step, getPipeShape(currentPos, mapLines))
		pipeLength++
	}
	return pipeLength/2 + 1
}

func getInsideNeighbour(clockwise bool, prevStep Step, pos Position, shape rune) []Position {
	switch {
	case shape == '|' && clockwise && prevStep.y == -1:
		return []Position{{x: pos.x + 1, y: pos.y}}
	case shape == '|' && !clockwise && prevStep.y == -1:
		return []Position{{x: pos.x - 1, y: pos.y}}
	case shape == '|' && clockwise && prevStep.y == 1:
		return []Position{{x: pos.x - 1, y: pos.y}}
	case shape == '|' && !clockwise && prevStep.y == 1:
		return []Position{{x: pos.x + 1, y: pos.y}}

	case shape == '-' && clockwise && prevStep.x == -1:
		return []Position{{x: pos.x, y: pos.y - 1}}
	case shape == '-' && !clockwise && prevStep.x == -1:
		return []Position{{x: pos.x, y: pos.y + 1}}
	case shape == '-' && clockwise && prevStep.x == 1:
		return []Position{{x: pos.x, y: pos.y + 1}}
	case shape == '-' && !clockwise && prevStep.x == 1:
		return []Position{{x: pos.x, y: pos.y - 1}}

	case shape == '7' && clockwise && prevStep.y == -1:
		return []Position{{x: pos.x + 1, y: pos.y}, {x: pos.x, y: pos.y - 1}}
	case shape == '7' && !clockwise && prevStep.y != -1:
		return []Position{{x: pos.x + 1, y: pos.y}, {x: pos.x, y: pos.y - 1}}
	case shape == 'J' && clockwise && prevStep.x == 1:
		return []Position{{x: pos.x + 1, y: pos.y}, {x: pos.x, y: pos.y + 1}}
	case shape == 'J' && !clockwise && prevStep.x != 1:
		return []Position{{x: pos.x + 1, y: pos.y}, {x: pos.x, y: pos.y + 1}}
	case shape == 'L' && clockwise && prevStep.y == 1:
		return []Position{{x: pos.x - 1, y: pos.y}, {x: pos.x, y: pos.y + 1}}
	case shape == 'L' && !clockwise && prevStep.y != 1:
		return []Position{{x: pos.x - 1, y: pos.y}, {x: pos.x, y: pos.y + 1}}
	case shape == 'F' && clockwise && prevStep.x == -1:
		return []Position{{x: pos.x - 1, y: pos.y}, {x: pos.x, y: pos.y - 1}}
	case shape == 'F' && !clockwise && prevStep.x != -1:
		return []Position{{x: pos.x - 1, y: pos.y}, {x: pos.x, y: pos.y - 1}}
	}
	return []Position{}
}

func ProblemTwo(scanner bufio.Scanner) int {
	mapLines := [][]rune{}
	for scanner.Scan() {
		mapLines = append(mapLines, []rune(scanner.Text()))
	}
	pipeLength := 0
	currentPos := findStart(mapLines)
	step := findFirstStep(mapLines, currentPos)
	netTurns := 0
	pipeParts := make(map[Position]bool)

	for {
		pipeParts[currentPos] = true
		currentPos.x += step.x
		currentPos.y += step.y
		shape := getPipeShape(currentPos, mapLines)
		if shape == 'S' {
			break
		}
		nextStep, turn := getNextStep(step, shape)
		step = nextStep
		netTurns += turn
		pipeLength++
	}

	clockwise := netTurns > 0

	enclosedSpaces := make(map[Position]bool)
	currentPos = findStart(mapLines)
	step = findFirstStep(mapLines, currentPos)
	for {
		currentPos.x += step.x
		currentPos.y += step.y
		shape := getPipeShape(currentPos, mapLines)
		if shape == 'S' {
			break
		}
		innerNeighbours := getInsideNeighbour(clockwise, step, currentPos, shape)
		for _, n := range innerNeighbours {
			enclosedSpaces[n] = true
		}
		nextStep, _ := getNextStep(step, shape)
		step = nextStep
	}

	maps.DeleteFunc(enclosedSpaces, func(pos Position, _ bool) bool { return pipeParts[pos] })
	for pos := range enclosedSpaces {
		scanPos := pos
		for {
			scanPos.x++
			if pipeParts[scanPos] {
				break
			}
			enclosedSpaces[scanPos] = true
		}
	}

	for y, line := range mapLines {
		for x, v := range line {
			switch {
			case enclosedSpaces[Position{x: x, y: y}]:
				fmt.Print("x")
			case pipeParts[Position{x, y}]:
				switch v {
				case 'S':
					fmt.Print("★")
				case '|':
					fmt.Print("║")
				case '-':
					fmt.Print("═")
				case 'F':
					fmt.Print("╔")
				case '7':
					fmt.Print("╗")
				case 'J':
					fmt.Print("╝")
				case 'L':
					fmt.Print("╚")
				}
			default:
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}

	return len(enclosedSpaces)
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
