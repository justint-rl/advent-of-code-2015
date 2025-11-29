package main

import (
	"fmt"

	wrappingpaper "github.com/justint-rl/advent-of-code-2015/day-2/go/wrapping_paper"
	fileprocessor "github.com/justint-rl/advent-of-code-2015/go_libs/file_processor"
)

func main() {
	fp := fileprocessor.New()
	data := fp.LoadLines("day-2/go/input/input.txt")
	wp := wrappingpaper.New()

	// Part 1
	result1 := wp.Part1(data)
	fmt.Printf("Part 1: %d\n", result1)

	// Part 2
	result2 := wp.Part2(data)
	fmt.Printf("Part 2: %d\n", result2)
}
