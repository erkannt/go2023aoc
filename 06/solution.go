package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	duration int
	record   int
}

func parseRaces(scanner bufio.Scanner) []Race {
	numberRegex, _ := regexp.Compile("[0-9]+")
	races := []Race{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Time: ") {
			numbers := numberRegex.FindAllString(line, -1)
			for _, number := range numbers {
				value, _ := strconv.Atoi(number)
				races = append(races, Race{duration: value})
			}
		}
		if strings.HasPrefix(line, "Distance: ") {
			numbers := numberRegex.FindAllString(line, -1)
			for i, number := range numbers {
				value, _ := strconv.Atoi(number)
				races[i].record = value
			}
		}
	}
	return races
}

func parseRacesIgnoringBadKerning(scanner bufio.Scanner) Race {
	numberRegex, _ := regexp.Compile("[0-9]+")
	race := Race{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, " ", "")
		if strings.HasPrefix(line, "Time:") {
			number := numberRegex.FindString(line)
			value, _ := strconv.Atoi(number)
			race.duration = value
		}
		if strings.HasPrefix(line, "Distance:") {
			number := numberRegex.FindString(line)
			value, _ := strconv.Atoi(number)
			race.record = value
		}
	}
	return race
}

func optimiseForRace(race Race) (int, int) {
	durationSquared := math.Pow(float64(race.duration), 2)
	minButton := math.Floor(float64(race.duration)/2 - math.Sqrt(durationSquared/4-float64(race.record)))
	maxButton := math.Ceil(float64(race.duration)/2+math.Sqrt(durationSquared/4-float64(race.record))) - 1
	return int(minButton), int(maxButton)
}

func ProblemOne(scanner bufio.Scanner) int {
	total := 1
	races := parseRaces(scanner)
	for _, race := range races {
		minButton, maxButton := optimiseForRace(race)
		total *= maxButton - minButton
	}
	return total
}

func ProblemTwo(scanner bufio.Scanner) int {
	race := parseRacesIgnoringBadKerning(scanner)
	minButton, maxButton := optimiseForRace(race)
	return maxButton - minButton
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
