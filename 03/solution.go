package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Location struct {
	x int
	y int
}

type Partnumber struct {
	value    int
	location Location
	lenght   int
}

func parseSchematic(scanner bufio.Scanner) ([]Partnumber, []Location) {
	symbolRegx, _ := regexp.Compile("[^0-9.]")
	numberRegx, _ := regexp.Compile("[0-9]+")

	var numbers = []Partnumber{}
	var locations = []Location{}

	var y = 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		symbolLocs := symbolRegx.FindAllIndex([]byte(line), -1)
		for _, loc := range symbolLocs {
			locations = append(locations, Location{x: loc[0], y: y})
		}

		numberLocs := numberRegx.FindAllIndex([]byte(line), -1)
		for _, loc := range numberLocs {
			value, _ := strconv.Atoi(line[loc[0]:loc[1]])
			newNumber := Partnumber{
				value:  value,
				lenght: loc[1] - loc[0],
				location: Location{
					x: loc[0],
					y: y,
				},
			}
			numbers = append(numbers, newNumber)
		}
		y += 1
	}

	return numbers, locations
}

func isAdjacent(number Partnumber, locations []Location) bool {
	for _, loc := range locations {
		if loc.x >= number.location.x-1 && loc.x <= number.location.x+number.lenght && loc.y >= number.location.y-1 && loc.y <= number.location.y+1 {
			return true
		}
	}
	return false
}

func ProblemOne(scanner bufio.Scanner) int {
	var total = 0
	partNumbers, symbolLocations := parseSchematic(scanner)
	for _, number := range partNumbers {
		if isAdjacent(number, symbolLocations) {
			total += number.value
		}
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
