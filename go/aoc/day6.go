package aoc

import (
	"log"
	"strings"

	"github.com/stangles/advent-of-code-2023/util"
)

type race struct {
	time     int
	distance int
}

func getRaceInput() []race {
	lines, err := util.GetStringInput("data/day6.txt")
	if err != nil {
		log.Fatalf("unable to get data: %v", err)
	}

	times := getEntryInts(lines[0])
	distances := getEntryInts(lines[1])

	races := make([]race, len(times))
	for i, time := range times {
		races[i] = race{time, distances[i]}
	}

	return races
}

func getEntryInts(entry string) []int {
	split := strings.Split(strings.TrimSpace(strings.Split(entry, ":")[1]), " ")
	nums := make([]int, 0)
	for _, el := range split {
		if el != "" {
			nums = append(nums, mustAtoi(el))
		}
	}

	return nums
}

func Day6Part1() int {
	races := getRaceInput()
	numWays := 1
	for _, race := range races {
		numWays *= simulate(race.time, race.distance)
	}
	return numWays
}

func simulate(time, recordDistance int) int {
	ways := 0
	for i := 1; i < time; i++ {
		if dist := i * (time - i); dist > recordDistance {
			ways++
		}
	}
	return ways
}

func Day6Part2() int {
	// turns out the solution to part 1 was fast enough if you just manually adjust the input
	return 0
}
