package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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

func getMinCubes(input game) reveal {
	return reveal{
		blue:  0,
		red:   0,
		green: 0,
	}
}

func ProblemTwo(scanner bufio.Scanner) int {
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		game := toGame(line)
		minCubes := getMinCubes(game)
		total += minCubes.red * minCubes.green * minCubes.blue
	}
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
