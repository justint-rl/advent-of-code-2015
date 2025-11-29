package solution

import (
	"fmt"

	constants "github.com/justint-rl/advent-of-code-2015/day-3/go/constants"
)

type HouseVisits interface {
	Part1(route string) int64
}

type houseVisitsImpl struct{}

func New() HouseVisits {
	return &houseVisitsImpl{}
}

func (h *houseVisitsImpl) Part1(route string) int64 {
	visits := int64(0)
	for _, direction := range route {
		switch constants.Direction(direction) {
		case constants.North:
			fmt.Println("Went north")
		case constants.South:
			fmt.Println("Went south")
		case constants.West:
			fmt.Println("Went west")
		case constants.East:
			fmt.Println("Went east")
		}
	}
	return visits
}
