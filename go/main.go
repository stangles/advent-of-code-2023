package main

import (
	"fmt"

	"github.com/stangles/advent-of-code-2023/aoc"
	"github.com/stangles/advent-of-code-2023/util"
)

func main() {
	ret := util.WithTimings(func() interface{} {
		return aoc.Day6Part1()
	})
	fmt.Println(ret)
}
