package aoc

import (
	"log"
	"math"
	"strings"

	"github.com/stangles/advent-of-code-2023/util"
)

type card struct {
	winning map[int]struct{}
	have    map[int]struct{}
}

func getCardInput() map[int]card {
	lines, err := util.GetStringInput("data/day4.txt")
	if err != nil {
		log.Fatalf("unable to get data: %v", err)
	}

	cards := make(map[int]card, len(lines))
	for _, line := range lines {
		cardIdSplit := strings.Split(line, ": ")
		cardId := mustAtoi(strings.TrimSpace(strings.Split(cardIdSplit[0], "Card ")[1]))

		numSplit := strings.Split(cardIdSplit[1], " | ")
		winStr := strings.Split(strings.TrimSpace(numSplit[0]), " ")
		haveStr := strings.Split(strings.TrimSpace(numSplit[1]), " ")

		winNums := toSet(strToInt(winStr))
		haveNums := toSet(strToInt(haveStr))
		cards[cardId] = card{
			winNums,
			haveNums,
		}
	}

	return cards
}

func strToInt(intsStr []string) []int {
	ints := make([]int, 0)
	for _, num := range intsStr {
		if num != "" {
			ints = append(ints, mustAtoi(strings.TrimSpace(num)))
		}
	}

	return ints
}

func toSet(arr []int) map[int]struct{} {
	set := make(map[int]struct{}, len(arr))
	for _, el := range arr {
		set[el] = struct{}{}
	}

	return set
}

func Day4Part1() int {
	cards := getCardInput()
	sum := 0
	for _, card := range cards {
		sum += int(math.Pow(2, float64(intersectLen(card.winning, card.have)-1)))
	}
	return sum
}

func intersectLen(set1, set2 map[int]struct{}) int {
	len := 0
	for k := range set1 {
		if _, ok := set2[k]; ok {
			len++
		}
	}

	return len
}

func Day4Part2() int {
	cards := getCardInput()
	counts := make(map[int]int, len(cards)+1)
	for i := 1; i < len(cards)+1; i++ {
		counts[i] += 1
		card := cards[i]
		matches := intersectLen(card.winning, card.have)
		for j := i + 1; j <= i+matches; j++ {
			counts[j] += counts[i]
		}
	}

	sum := 0
	for _, count := range counts {
		sum += count
	}
	return sum
}
