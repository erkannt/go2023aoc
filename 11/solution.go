package main

import (
	"log"
	"math"
	"os"
	"slices"
	"strings"
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

func findEmptyColumns(image Image) []int {
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
	return emptyColumns
}

func findEmptyRows(image Image) []int {
	emptyLines := []int{}
	for i, line := range image {
		if !slices.Contains(line, '#') {
			emptyLines = append(emptyLines, i)
		}
	}
	return emptyLines
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

func measureDistances(positions []Position, emptyRows []int, emptyColumns []int, expansionFactor int) []int {
	galaxyCount := len(positions)
	distances := []int{}
	for i := 0; i < galaxyCount; i++ {
		for j := i + 1; j < galaxyCount; j++ {
			galaxyA := positions[i]
			galaxyB := positions[j]
			colDist := math.Abs(float64(galaxyB.x - galaxyA.x))
			rowDist := galaxyB.y - galaxyA.y
			expansion := 0
			for _, row := range emptyRows {
				if row > galaxyA.y && row < galaxyB.y {
					expansion += expansionFactor - 1
				}
			}
			for _, col := range emptyColumns {
				switch galaxyA.x < galaxyB.x {
				case true:
					if col > galaxyA.x && col < galaxyB.x {
						expansion += expansionFactor - 1
					}
				case false:
					if col < galaxyA.x && col > galaxyB.x {
						expansion += expansionFactor - 1
					}
				}
			}
			distances = append(distances, int(colDist)+rowDist+expansion)
		}
	}
	return distances
}

func ProblemOne(input string) int {
	return ProblemTwo(input, 2)
}

func ProblemTwo(input string, expansion int) int {
	image := parseImage(input)
	emptyRows := findEmptyRows(image)
	emptyColumns := findEmptyColumns(image)
	galaxies := findGalaxies(image)
	distances := measureDistances(galaxies, emptyRows, emptyColumns, expansion)
	total := 0
	for _, d := range distances {
		total += d
	}
	return total
}

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	println("Problem1:", ProblemOne(string(input)))
	println("Problem2:", ProblemTwo(string(input), 1000000))
}
