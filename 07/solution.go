package main

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"regexp"
	"slices"
	"sort"
	"strconv"
)

type Hand struct {
	cards string
	bid   int
}

type CardsType int

const (
	HighCard CardsType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func getType(cards string) CardsType {
	uniqueCards := map[rune]int{}
	for _, v := range cards {
		uniqueCards[v] += 1
	}
	counts := []int{}
	for _, v := range uniqueCards {
		counts = append(counts, v)
	}
	slices.Sort(counts)
	slices.Reverse(counts)
	switch {
	case reflect.DeepEqual(counts, []int{5}):
		return FiveOfAKind
	case reflect.DeepEqual(counts, []int{4, 1}):
		return FourOfAKind
	case reflect.DeepEqual(counts, []int{3, 2}):
		return FullHouse
	case reflect.DeepEqual(counts, []int{3, 1, 1}):
		return ThreeOfAKind
	case reflect.DeepEqual(counts, []int{2, 2, 1}):
		return TwoPair
	case reflect.DeepEqual(counts, []int{2, 1, 1, 1}):
		return OnePair
	default:
		return HighCard
	}
}

func getTypeWithJoker(cards string) CardsType {
	uniqueCards := map[rune]int{}
	for _, v := range cards {
		uniqueCards[v] += 1
	}
	counts := []int{}
	for _, v := range uniqueCards {
		counts = append(counts, v)
	}
	slices.Sort(counts)
	slices.Reverse(counts)
	switch {
	case reflect.DeepEqual(counts, []int{5}):
		return FiveOfAKind
	case reflect.DeepEqual(counts, []int{4, 1}):
		switch uniqueCards['J'] {
		case 4:
			return FiveOfAKind
		case 1:
			return FiveOfAKind
		default:
			return FourOfAKind
		}
	case reflect.DeepEqual(counts, []int{3, 2}):
		switch uniqueCards['J'] {
		case 3:
			return FiveOfAKind
		case 2:
			return FiveOfAKind
		default:
			return FullHouse
		}
	case reflect.DeepEqual(counts, []int{3, 1, 1}):
		switch uniqueCards['J'] {
		case 3:
			return FourOfAKind
		case 1:
			return FourOfAKind
		default:
			return ThreeOfAKind
		}
	case reflect.DeepEqual(counts, []int{2, 2, 1}):
		switch uniqueCards['J'] {
		case 2:
			return FourOfAKind
		case 1:
			return FullHouse
		default:
			return TwoPair
		}
	case reflect.DeepEqual(counts, []int{2, 1, 1, 1}):
		switch uniqueCards['J'] {
		case 2:
			return ThreeOfAKind
		case 1:
			return ThreeOfAKind
		default:
			return OnePair
		}
	default:
		switch uniqueCards['J'] {
		case 1:
			return OnePair
		default:
			return HighCard
		}
	}
}

const partOneValues = "23456789TJQKA"

const partTwoValues = "J23456789TQKA"

func cardValue(input byte, cardsByValue string) int {
	for i := 0; i < len(cardsByValue); i++ {
		if cardsByValue[i] == input {
			return i
		}
	}
	return -1
}

func firstDifferentCardIsLower(cardsA, cardsB, cardsByValue string) bool {
	for i := 0; i <= 4; i++ {
		valueA := cardValue(cardsA[i], cardsByValue)
		valueB := cardValue(cardsB[i], cardsByValue)
		if valueA > valueB {
			return false
		}
		if valueA < valueB {
			return true
		}
	}
	return false
}

type ByRank []Hand

func (a ByRank) Len() int      { return len(a) }
func (a ByRank) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRank) Less(i, j int) bool {
	cardsI := a[i].cards
	cardsJ := a[j].cards
	cardsTypeI := getType(cardsI)
	cardsTypeJ := getType(cardsJ)
	switch {
	case cardsTypeI < cardsTypeJ:
		return true
	case cardsTypeI > cardsTypeJ:
		return false
	default:
		return firstDifferentCardIsLower(cardsI, cardsJ, partOneValues)
	}
}

type ByRankWithJoker []Hand

func (a ByRankWithJoker) Len() int      { return len(a) }
func (a ByRankWithJoker) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRankWithJoker) Less(i, j int) bool {
	cardsI := a[i].cards
	cardsJ := a[j].cards
	cardsTypeI := getTypeWithJoker(cardsI)
	cardsTypeJ := getTypeWithJoker(cardsJ)
	switch {
	case cardsTypeI < cardsTypeJ:
		return true
	case cardsTypeI > cardsTypeJ:
		return false
	default:
		return firstDifferentCardIsLower(cardsI, cardsJ, partTwoValues)
	}
}

func parseHands(scanner bufio.Scanner) []Hand {
	handRegex, _ := regexp.Compile("([AKQJT98765432]{5}) ([0-9]+)")
	hands := []Hand{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		groups := handRegex.FindAllStringSubmatch(line, -1)
		cards := groups[0][1]
		value, _ := strconv.Atoi(groups[0][2])
		hands = append(hands, Hand{
			cards: cards,
			bid:   value,
		})
	}
	return hands
}

func tally(hands []Hand) int {
	total := 0
	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}
	return total
}

func ProblemOne(scanner bufio.Scanner) int {
	hands := parseHands(scanner)
	sort.Sort(ByRank(hands))
	return tally(hands)
}

func ProblemTwo(scanner bufio.Scanner) int {
	hands := parseHands(scanner)
	sort.Sort(ByRankWithJoker(hands))
	return tally(hands)
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
