package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Seed = int

type SeedMapping struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}

func parseAlmanac(scanner bufio.Scanner) ([]Seed, [][]SeedMapping) {
	numberRegex, _ := regexp.Compile("[0-9]+")
	seeds := []Seed{}
	allMaps := [][]SeedMapping{}
	currentMap := []SeedMapping{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.Contains(line, "seeds: ") {
			seedStrings := numberRegex.FindAllString(line, -1)
			for _, str := range seedStrings {
				value, _ := strconv.Atoi(str)
				seeds = append(seeds, Seed(value))
			}
			continue
		}

		if strings.Contains(line, "map:") && len(currentMap) != 0 {
			allMaps = append(allMaps, currentMap)
			currentMap = []SeedMapping{}
			continue
		}

		numberStrings := numberRegex.FindAllString(line, -1)
		if len(numberStrings) == 3 {
			numbers := []int{}
			for _, str := range numberStrings {
				value, _ := strconv.Atoi(str)
				numbers = append(numbers, value)
			}
			currentMap = append(currentMap, SeedMapping{
				destinationStart: numbers[0],
				sourceStart:      numbers[1],
				rangeLength:      numbers[2],
			})
		}
	}
	allMaps = append(allMaps, currentMap)

	return seeds, allMaps
}

func lookupLocation(seed Seed, maps [][]SeedMapping) int {
	var result = seed
	for _, seedMap := range maps {
		for _, mapping := range seedMap {
			if result >= mapping.sourceStart && result <= mapping.sourceStart+mapping.rangeLength {
				delta := result - mapping.sourceStart
				result = mapping.destinationStart + delta
				break
			}
		}
	}
	return result
}

func ProblemOne(scanner bufio.Scanner) int {
	seeds, maps := parseAlmanac(scanner)
	var nearestLocation = lookupLocation(seeds[0], maps)
	for _, seed := range seeds {
		location := lookupLocation(seed, maps)
		if location < nearestLocation {
			nearestLocation = location
		}
	}
	return nearestLocation
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	println("Problem1:", ProblemOne(*bufio.NewScanner(file)))
}
