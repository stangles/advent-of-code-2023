package aoc

import (
	"log"
	"math"
	"strings"
	"unicode"

	"github.com/stangles/advent-of-code-2023/util"
)

func getAlmanac() ([]int, []*node) {
	lines, err := util.GetStringInput("data/day5.txt")
	if err != nil {
		log.Fatalf("unable to get data: %v", err)
	}

	seedsStr := strings.Split(strings.Split(lines[0], "seeds: ")[1], " ")
	seeds := make([]int, len(seedsStr))
	for i := 0; i < len(seedsStr); i++ {
		seeds[i] = mustAtoi(seedsStr[i])
	}

	mappingNodes := make([]*node, 7)
	nodeIdx := -1
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			nodeIdx++
			continue
		}
		if unicode.IsDigit([]rune(line)[0]) {
			mapStr := strings.Split(line, " ")
			destRangeStart := mustAtoi(mapStr[0])
			sourceRangeStart := mustAtoi(mapStr[1])
			rangeTot := mustAtoi(mapStr[2])

			n := &node{
				sourceRangeStart,
				sourceRangeStart + rangeTot - 1,
				destRangeStart - sourceRangeStart,
				nil,
				nil,
			}
			if root := mappingNodes[nodeIdx]; root != nil {
				root.insert(n)
			} else {
				mappingNodes[nodeIdx] = n
			}
		}
	}

	return seeds, mappingNodes
}

type node struct {
	start int
	end   int
	adj   int
	left  *node
	right *node
}

func (root *node) insert(n *node) {
	if n.start < root.start {
		if l := root.left; l != nil {
			l.insert(n)
		} else {
			root.left = n
		}
	} else if n.start > root.start {
		if r := root.right; r != nil {
			r.insert(n)
		} else {
			root.right = n
		}
	}
}

func Day5Part1() int {
	seeds, nodes := getAlmanac()
	lowest := math.MaxInt
	for _, seed := range seeds {
		mapped := seed
		for _, n := range nodes {
			mapped = convert(mapped, n)
		}
		if mapped < lowest {
			lowest = mapped
		}
	}

	return lowest
}

func convert(in int, mappingNode *node) int {
	cur := mappingNode
	for cur != nil {
		if in < cur.start {
			cur = cur.left
		} else if in >= cur.start {
			if in <= cur.end {
				return in + cur.adj
			} else {
				cur = cur.right
			}
		}
	}

	return in
}

func Day5Part2() int {
	seeds, nodes := getAlmanac()
	lowest := math.MaxInt
	for i := 0; i < len(seeds); i += 2 {
		for start := seeds[i]; start < seeds[i]+seeds[i+1]; start++ {
			mapped := start
			for _, n := range nodes {
				mapped = convert(mapped, n)
			}
			if mapped < lowest {
				lowest = mapped
			}
		}
	}

	return lowest
}
