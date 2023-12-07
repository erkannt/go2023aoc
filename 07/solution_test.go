package main

import (
	"bufio"
	"reflect"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"testing"
)

var problemInput = `
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

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

var cardsByValue = []string{
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"T",
	"J",
	"Q",
	"K",
	"A",
}

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

func cardValue(input byte) int {
	cardsByValue := "23456789TJQKA"
	for i := 0; i < len(cardsByValue); i++ {
		if cardsByValue[i] == input {
			return i
		}
	}
	return -1
}

func firstDifferentCardIsLower(cardsA, cardsB string) bool {
	for i := 0; i <= 4; i++ {
		valueA := cardValue(cardsA[i])
		valueB := cardValue(cardsB[i])
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
	cardsTypeI := getType(a[i].cards)
	cardsTypeJ := getType(a[j].cards)
	if cardsTypeI < cardsTypeJ {
		return true
	}
	return firstDifferentCardIsLower(a[i].cards, a[j].cards)
}

func ProblemOne(scanner bufio.Scanner) int {
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
	sort.Sort(ByRank(hands))
	total := 0
	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}
	return total
}

func TestProblemOne(t *testing.T) {
	expected := 6440

	result := ProblemOne(*bufio.NewScanner(strings.NewReader(problemInput)))
	if result != expected {
		t.Errorf("\n Expected: %v\nResult: %v", expected, result)
	}
}
