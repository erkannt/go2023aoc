package main

import (
	"fmt"
	"math"
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
	emptyColumns := []int{}
	width := len(image[0])
	height := len(image)
	for col := 0; col < width; col++ {
		empty := true
		for row := 0; row < height; row++ {
			if image[row][col] == '#' {
				empty = false
				break
			}
		}
		if empty {
			emptyColumns = append(emptyColumns, col)
		}
	}
	newImage := Image{}
	for _, row := range image {
		for i, col := range emptyColumns {
			row = slices.Insert(row, col+i, '.')
		}
		newImage = append(newImage, row)
	}
	return newImage
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
	galaxies := []Position{}
	for rowNo, row := range image {
		for colNo, value := range row {
			if value == '#' {
				galaxies = append(galaxies, Position{x: colNo, y: rowNo})
			}
		}
	}
	return galaxies
}

func measureDistances(positions []Position) []int {
	galaxyCount := len(positions)
	distances := []int{}
	for i := 0; i < galaxyCount; i++ {
		for j := i + 1; j < galaxyCount; j++ {
			galaxyA := positions[i]
			galaxyB := positions[j]
			dist := math.Abs(float64(galaxyA.x)-float64(galaxyB.x)) + math.Abs(float64(galaxyA.y)-float64(galaxyB.y))
			distances = append(distances, int(dist))
		}
	}
	return distances
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
	println("After horizontal expansion")
	printImage(image)
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
