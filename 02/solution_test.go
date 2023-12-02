package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

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

func TestProblemTwo(t *testing.T) {
	input := `
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`

	reader := bufio.NewScanner(strings.NewReader(input))
	ans := ProblemTwo(*reader)
	if ans != 2286 {
		t.Error("wrong result", ans)
	}
}

func TestToGame(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	expected := game{
		id: 1,
		reveals: []cubeset{
			{
				blue: 3,
				red:  4,
			},
			{
				red:   1,
				green: 2,
				blue:  6,
			},
			{
				green: 2,
			},
		},
	}

	ans := toGame(input)
	if !reflect.DeepEqual(ans, expected) {
		t.Error("wrong result", ans)
	}
}
