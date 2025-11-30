package main

import (
	"fmt"
	"os"

	housevisits "github.com/justint-rl/advent-of-code-2015/day-3/go/solution"
	fileprocessor "github.com/justint-rl/advent-of-code-2015/go_libs/file_processor"
)

func main() {
	fmt.Println("All args:", os.Args)

	if len(os.Args) < 2 {
		fmt.Println("Missing input file path")
		return
	}

	filePath := os.Args[1]

	fp := fileprocessor.New()
	data := fp.LoadRaw(filePath)
	hv := housevisits.New()

	// Part 1
	result1 := hv.Part1(data)
	fmt.Printf("Part 1: %d\n", result1)

	// Part 2
	result2 := hv.Part2(data)
	fmt.Printf("Part 2: %d\n", result2)
}
