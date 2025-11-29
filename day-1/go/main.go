package main

import (
	"fmt"

	notquitelisp "github.com/justint-rl/advent-of-code-2015/day-1/not_quite_lisp"
	fileprocessor "github.com/justint-rl/advent-of-code-2015/go_libs/file_processor"
)

func main() {
	fp := fileprocessor.New()
	data := fp.LoadRaw("day-1/go/input/input.txt")
	nql := notquitelisp.New()

	// Part 1
	result1 := nql.Part1(data)
	fmt.Printf("Part 1: %d\n", result1)

	// Part 2
	result2 := nql.Part2(data)
	fmt.Printf("Part 2: %d\n", result2)
}
