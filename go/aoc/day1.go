package aoc

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/stangles/advent-of-code-2023/util"
)

func Day1Part1() int {
	lines, err := util.GetStringInput("data/day1.txt")
	if err != nil {
		log.Fatalf("unable to get data: %v", err)
	}

	sum := 0
	for _, line := range lines {
		calVal := ""
		prevVal := -1
		for i, c := range line {
			if val, err := strconv.Atoi(string(c)); err == nil {
				if prevVal == -1 {
					calVal += fmt.Sprint(val)
				}
				prevVal = val
			}
			if i+1 == len([]rune(line)) {
				calVal += fmt.Sprint(prevVal)
			}
		}
		calValInt, err := strconv.Atoi(calVal)
		if err != nil {
			panic(err)
		}

		sum += calValInt
	}

	return sum
}

func Day1Part2() int {
	lines, err := util.GetStringInput("data/day1.txt")
	if err != nil {
		log.Fatalf("unable to get data: %v", err)
	}

	sum := 0
	for _, line := range lines {
		calVal := ""
		prevVal := -1

		endPrefix := 0
		for i, c := range line {
			if unicode.IsDigit(c) {
				endPrefix = i
				break
			}
		}

		transformedPrefix := transformWordToInt(line[0:endPrefix], false)

		runeS := []rune(line)
		startSuffix := len(runeS)
		for i := len(runeS) - 1; i >= 0; i-- {
			if unicode.IsDigit(runeS[i]) {
				startSuffix = i + 1
				break
			}
		}

		transformedSuffix := transformWordToInt(line[startSuffix:len(runeS)], true)

		transformed := transformedPrefix + line[endPrefix:startSuffix] + transformedSuffix
		for i, c := range transformed {
			if val, err := strconv.Atoi(string(c)); err == nil {
				if prevVal == -1 {
					calVal += fmt.Sprint(val)
				}
				prevVal = val
			}
			if i+1 == len([]rune(transformed)) {
				calVal += fmt.Sprint(prevVal)
			}
		}
		calValInt, err := strconv.Atoi(calVal)
		if err != nil {
			panic(err)
		}

		sum += calValInt
	}

	return sum
}

func transformWordToInt(s string, reverse bool) string {
	wordToInt := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	words := make([]string, len(wordToInt))
	i := 0
	for k := range wordToInt {
		words[i] = k
		i++
	}

	if reverse {
		highestFoundIdx := -1
		foundHigh := ""
		for i := 0; i < len(words); i++ {
			foundHighIdx := strings.LastIndex(s, words[i])
			if foundHighIdx != -1 && foundHighIdx > highestFoundIdx {
				foundHigh = words[i]
				highestFoundIdx = foundHighIdx
			}
		}
		if highestFoundIdx > -1 {
			s = s[:highestFoundIdx] + fmt.Sprint(wordToInt[foundHigh]) + s[highestFoundIdx+1:]
		}
	} else {
		lowestFoundIdx := len([]rune(s))
		foundLow := ""
		for i := 0; i < len(words); i++ {
			foundLowIdx := strings.Index(s, words[i])
			if foundLowIdx != -1 && foundLowIdx < lowestFoundIdx {
				foundLow = words[i]
				lowestFoundIdx = foundLowIdx
			}
		}
		if lowestFoundIdx < len([]rune(s)) {
			s = strings.Replace(s, foundLow, fmt.Sprint(wordToInt[foundLow]), 1)
		}
	}

	return s
}
