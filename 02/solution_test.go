package main

import (
	"bufio"
	"log"
	"reflect"
	"strconv"
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

func toId(input string) int {
	idString, _ := strings.CutPrefix(input, "Game ")
	idValue, err := strconv.Atoi(idString)
	if err != nil {
		log.Fatalf("Failed to parse game id from: %s", input)
	}
	return idValue
}

func toReveals(input string) []reveal {
	revealStrings := strings.Split(input, "; ")
	var reveals = []reveal{}
	for _, revealString := range revealStrings {
		colorParts := strings.Split(revealString, ", ")
		var reveal = reveal{}
		for _, colorPart := range colorParts {
			colorPart = strings.Trim(colorPart, " ")
			countString, colorName, _ := strings.Cut(colorPart, " ")
			countValue, err := strconv.Atoi(countString)
			if err != nil {
				log.Fatalf("Failed to parse color count from: '%s' (part of: '%s')", countString, colorPart)
			}
			switch colorName {
			case "red":
				reveal.red = countValue
			case "green":
				reveal.green = countValue
			case "blue":
				reveal.blue = countValue
			}
		}
		reveals = append(reveals, reveal)
	}
	return reveals

}

func toGame(input string) game {
	idPart, revealsPart, _ := strings.Cut(input, ":")

	return game{
		id:      toId(idPart),
		reveals: toReveals(revealsPart),
	}
}

func isPossible(game game) bool {
	for _, reveal := range game.reveals {
		if reveal.red > 12 || reveal.green > 13 || reveal.blue > 14 {
			return false
		}
	}
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

func TestToGame(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	expected := game{
		id: 1,
		reveals: []reveal{
			reveal{
				blue: 3,
				red:  4,
			},
			reveal{
				red:   1,
				green: 2,
				blue:  6,
			},
			reveal{
				green: 2,
			},
		},
	}

	ans := toGame(input)
	if !reflect.DeepEqual(ans, expected) {
		t.Error("wrong result", ans)
	}
}
