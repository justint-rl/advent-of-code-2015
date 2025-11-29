package main

import (
	"fmt"

	fileprocessor "github.com/justint-rl/advent-of-code-2015/go_libs/file_processor"
)

func main() {
	fp := fileprocessor.New()
	data := fp.LoadRaw("day-1/go/input/input.txt")
	fmt.Println(data)
}
