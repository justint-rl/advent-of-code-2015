package main

import (
	"fmt"

	housevisits "github.com/justint-rl/advent-of-code-2015/day-3/go/solution"
	fileprocessor "github.com/justint-rl/advent-of-code-2015/go_libs/file_processor"
)

func main() {
	fp := fileprocessor.New()
	data := fp.LoadRaw("day-3/go/input/input.txt")
	hv := housevisits.New()

	// Part 1
	result1 := hv.Part1(data)
	fmt.Printf("Part 1: %d\n", result1)

	// Part 2
	// result2 := hv.Part2(data)
	// fmt.Printf("Part 2: %d\n", result2)
}
