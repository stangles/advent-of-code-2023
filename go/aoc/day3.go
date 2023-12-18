package aoc

import (
	"log"
	"unicode"

	"github.com/stangles/advent-of-code-2023/util"
)

func getInput() [][]rune {
	lines, err := util.GetRuneInput("data/day3.txt")
	if err != nil {
		log.Fatalf("unable to get data: %v", err)
	}

	return lines
}

func Day3Part1() int {
	engine := getInput()
	sum := 0
	for i, line := range engine {
		numStartIdx := -1
		for j := 0; j < len(line); j++ {
			c := engine[i][j]
			if unicode.IsDigit(c) {
				if numStartIdx < 0 {
					numStartIdx = j
				}
				if matchR, _ := getAdjacentMatch(engine, i, j, func(row int, col int) bool {
					return !unicode.IsDigit(engine[row][col]) && engine[row][col] != '.'
				}); matchR != -1 {
					strPartNum := ""
					for k := numStartIdx; k < len(engine[i]) && unicode.IsDigit(engine[i][k]); k++ {
						strPartNum += string(engine[i][k])
						j = k
					}
					partNum := mustAtoi(strPartNum)
					sum += partNum
				}
			} else {
				numStartIdx = -1
			}
		}
	}
	return sum
}

func getAdjacentMatch(engine [][]rune, row int, col int, matchExpr func(int, int) bool) (int, int) {
	// check 3 above
	if row > 0 {
		rowAbove := row - 1
		if col > 0 {
			if matchExpr(rowAbove, col-1) {
				return rowAbove, col - 1
			}
		}
		if matchExpr(rowAbove, col) {
			return rowAbove, col
		}
		if col < len(engine[rowAbove])-1 {
			if matchExpr(rowAbove, col+1) {
				return rowAbove, col + 1
			}
		}
	}
	// check left and right
	if col > 0 {
		if matchExpr(row, col-1) {
			return row, col - 1
		}
	}
	if col < len(engine[row])-1 {
		if matchExpr(row, col+1) {
			return row, col + 1
		}
	}
	// check 3 below
	if row < len(engine)-1 {
		rowBelow := row + 1
		if col > 0 {
			if matchExpr(rowBelow, col-1) {
				return rowBelow, col - 1
			}
		}
		if matchExpr(rowBelow, col) {
			return rowBelow, col
		}
		if col < len(engine[rowBelow])-1 {
			if matchExpr(rowBelow, col+1) {
				return rowBelow, col + 1
			}
		}
	}
	return -1, -1
}

type gearCoord struct {
	row int
	col int
}

func Day3Part2() int {
	engine := getInput()
	gearCoordMap := make(map[gearCoord][]int, 0)

	for i, line := range engine {
		numStartIdx := -1
		for j := 0; j < len(line); j++ {
			c := engine[i][j]
			if unicode.IsDigit(c) {
				if numStartIdx < 0 {
					numStartIdx = j
				}
				if matchR, matchC := getAdjacentMatch(engine, i, j, func(row int, col int) bool {
					return engine[row][col] == '*'
				}); matchR != -1 {
					strPartNum := ""
					for k := numStartIdx; k < len(engine[i]) && unicode.IsDigit(engine[i][k]); k++ {
						strPartNum += string(engine[i][k])
						j = k
					}
					partNum := mustAtoi(strPartNum)
					gc := gearCoord{matchR, matchC}
					if _, ok := gearCoordMap[gc]; ok {
						gearCoordMap[gc] = append(gearCoordMap[gc], partNum)
					} else {
						gearCoordMap[gc] = []int{partNum}
					}
				}
			} else {
				numStartIdx = -1
			}
		}
	}

	ratioSum := 0
	for _, v := range gearCoordMap {
		if len(v) == 2 {
			ratioSum += v[0] * v[1]
		}
	}
	return ratioSum
}
