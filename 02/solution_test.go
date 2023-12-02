package main

import (
	"bufio"
	"strings"
	"testing"
)

type reveal struct {
	blue  int
	red   int
	green int
}

type game struct {
	id      int
	reveals []reveal
}

func toGame(input string) game {
	return game{
		id:      0,
		reveals: []reveal{},
	}
}

func isPossible(game game) bool {
	return true
}

func ProblemOne(scanner bufio.Scanner) int {
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		game := toGame(line)
		if isPossible(game) {
			total += game.id
		}
	}
	return total
}

func TestProblemOne(t *testing.T) {
	input := `
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`

	reader := bufio.NewScanner(strings.NewReader(input))
	ans := ProblemOne(*reader)
	if ans != 8 {
		t.Error("wrong result", ans)
	}
}
