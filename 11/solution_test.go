package main

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

type Image = [][]rune

type Position struct {
	x int
	y int
}

func parseImage(input string) Image {
	image := Image{}
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		image = append(image, []rune(l))
	}
	return image
}

func expandHorizontally(image Image) Image {
	return image
}

func expandVertically(image Image) Image {
	emptyLines := []int{}
	for i, line := range image {
		if !slices.Contains(line, '#') {
			emptyLines = append(emptyLines, i)
		}
	}
	for i, lineNo := range emptyLines {
		image = slices.Insert(image, lineNo+i, image[lineNo+i])
	}
	return image
}

func findGalaxies(image Image) []Position {
	return []Position{}
}

func measureDistances(positions []Position) []int {
	return []int{}
}

func printImage(image Image) {
	for _, v := range image {
		for _, c := range v {
			fmt.Printf("%v", string(c))
		}
		fmt.Printf("\n")
	}
}

func ProblemOne(input string) int {
	image := parseImage(input)
	printImage(image)
	image = expandVertically(image)
	println("After vertical expansion")
	printImage(image)
	image = expandHorizontally(image)
	galaxies := findGalaxies(image)
	distances := measureDistances(galaxies)
	total := 0
	for _, d := range distances {
		total += d
	}
	return total
}

func TestProblemOne(t *testing.T) {
	input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	expected := 374
	actual := ProblemOne(input)
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\v", expected, actual)
	}
}
