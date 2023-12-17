package aoc

import (
	"log"
	"strconv"
	"strings"

	"github.com/stangles/advent-of-code-2023/util"
)

type bagSet struct {
	red   int
	blue  int
	green int
}

func Day2Part1() int {
	lines, err := util.GetStringInput("data/day2.txt")
	if err != nil {
		log.Fatalf("unable to get data: %v", err)
	}

	games := make(map[int][]bagSet, len(lines))
	possible := 0
	for _, line := range lines {
		colon := strings.Split(line, ":")
		game := mustAtoi(strings.Split(colon[0], " ")[1])
		sets := strings.Split(colon[1], ";")
		// fmt.Println(fmt.Sprintf("game num %v", game))
		// fmt.Println(sets)
		games[game] = make([]bagSet, len(sets))

		isStillPossible := true
		for i, set := range sets {
			die := strings.Split(set, ",")
			bagSet := bagSet{}
			for _, dice := range die {
				tidy := strings.TrimSpace(dice)
				pair := strings.Split(tidy, " ")
				if pair[1] == "red" {
					bagSet.red = mustAtoi(pair[0])
					if isStillPossible {
						if bagSet.red > 12 {
							isStillPossible = false
						}
					}
				} else if pair[1] == "blue" {
					bagSet.blue = mustAtoi(pair[0])
					if isStillPossible {
						if bagSet.blue > 14 {
							isStillPossible = false
						}
					}
				} else {
					bagSet.green = mustAtoi(pair[0])
					if isStillPossible {
						if bagSet.green > 13 {
							isStillPossible = false
						}
					}
				}
			}
			games[game][i] = bagSet
		}
		if isStillPossible {
			possible += game
		}
	}
	// fmt.Println(games)

	return possible
}

func mustAtoi(intStr string) int {
	res, err := strconv.Atoi(intStr)
	if err != nil {
		panic(err)
	}

	return res
}

func Day2Part2() int {
	lines, err := util.GetStringInput("data/day2.txt")
	if err != nil {
		log.Fatalf("unable to get data: %v", err)
	}

	games := make(map[int][]bagSet, len(lines))
	for _, line := range lines {
		colon := strings.Split(line, ":")
		game := mustAtoi(strings.Split(colon[0], " ")[1])
		sets := strings.Split(colon[1], ";")
		games[game] = make([]bagSet, len(sets))

		for i, set := range sets {
			die := strings.Split(set, ",")
			bagSet := bagSet{}
			for _, dice := range die {
				tidy := strings.TrimSpace(dice)
				pair := strings.Split(tidy, " ")
				if pair[1] == "red" {
					bagSet.red = mustAtoi(pair[0])
				} else if pair[1] == "blue" {
					bagSet.blue = mustAtoi(pair[0])
				} else {
					bagSet.green = mustAtoi(pair[0])
				}
			}
			games[game][i] = bagSet
		}
	}

	powerSum := 0
	for _, sets := range games {
		minimums := []int{-1, -1, -1}
		for _, set := range sets {
			if set.red > minimums[0] {
				minimums[0] = set.red
			}
			if set.blue > minimums[1] {
				minimums[1] = set.blue
			}
			if set.green > minimums[2] {
				minimums[2] = set.green
			}
		}
		powerSum += minimums[0] * minimums[1] * minimums[2]
	}

	return powerSum
}
