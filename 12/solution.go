package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parse(input string) (string, []int) {
	numberRegex := regexp.MustCompile("[0-9]+")
	parts := strings.Split(input, " ")
	groupsStrings := numberRegex.FindAllStringSubmatch(parts[1], -1)
	groups := []int{}
	for _, v := range groupsStrings {
		value, _ := strconv.Atoi(v[0])
		groups = append(groups, value)
	}
	return parts[0], groups
}

func getCandidates(record string, groupSizes []int) [][]int {
	candidates := [][]int{}
	for _, size := range groupSizes {
		regx, _ := regexp.Compile(fmt.Sprintf("([?.][#?]{%d}[?.])", size))
		padded := fmt.Sprintf(".%s.", record)
		thisGroupsCandidates := []int{}
		offset := 0
		for {
			pos := regx.FindStringIndex(padded[offset:])
			if pos == nil {
				break
			}
			thisGroupsCandidates = append(thisGroupsCandidates, pos[0]+offset)
			offset += pos[0] + 1
		}
		if len(thisGroupsCandidates) == 0 {
			log.Fatalf("Can't find position for group: %v %v %v", record, groupSizes, size)
		}
		candidates = append(candidates, thisGroupsCandidates)
	}
	return candidates
}

type remainingGroupArrangements struct {
	record             string
	candidatePositions [][]int
	sizes              []int
}

func (r remainingGroupArrangements) calc(depth int, startingPoint int) int {
	if len(r.candidatePositions[depth:]) == 1 {
		possibleCount := 0
		for _, pos := range r.candidatePositions[depth:][0] {
			if pos >= startingPoint && !strings.ContainsRune(r.record[pos+r.sizes[depth:][0]:], '#') && !strings.ContainsRune(r.record[startingPoint:pos], '#') {
				possibleCount++
			}
		}
		return possibleCount
	}
	possibleCount := 0
	for _, pos := range r.candidatePositions[depth:][0] {
		if pos >= startingPoint && !strings.ContainsRune(r.record[startingPoint:pos], '#') {
			possibleCount += r.calc(depth+1, r.sizes[depth:][0]+pos+1)
		}
	}
	return possibleCount
}

func PossibleArrangements(input string, unfold bool) int {
	record, groupSizes := parse(input)
	if unfold {
		record = fmt.Sprintf("%s?%s?%s?%s?%s", record, record, record, record, record)
		expanded := []int{}
		for i := 0; i < 5; i++ {
			expanded = append(expanded, groupSizes...)
		}
		groupSizes = expanded
	}
	candidatePositions := getCandidates(record, groupSizes)
	calculator := remainingGroupArrangements{
		record:             record,
		candidatePositions: candidatePositions,
		sizes:              groupSizes,
	}
	total := calculator.calc(0, 0)
	return total
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	totalPart1 := 0
	totalPart2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		totalPart1 += PossibleArrangements(line, false)
		totalPart2 += PossibleArrangements(line, true)
	}

	println(totalPart1, totalPart2)
}
