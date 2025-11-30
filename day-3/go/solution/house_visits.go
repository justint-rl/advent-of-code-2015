package solution

import (
	constants "github.com/justint-rl/advent-of-code-2015/day-3/go/constants"
	model "github.com/justint-rl/advent-of-code-2015/day-3/go/model"
)

type HouseVisits interface {
	Part1(route string) int64
}

type houseVisitsImpl struct{}

func New() HouseVisits {
	return &houseVisitsImpl{}
}

func (h *houseVisitsImpl) Part1(route string) int64 {
	coordinatesVisited := make(map[model.Coordinate]struct{})
	currentCoordinate := model.New(0, 0)
	for _, direction := range route {
		coordinatesVisited[currentCoordinate] = struct{}{}
		switch constants.Direction(direction) {
		case constants.North:
			currentCoordinate = currentCoordinate.MoveUp(1)
		case constants.South:
			currentCoordinate = currentCoordinate.MoveDown(1)
		case constants.West:
			currentCoordinate = currentCoordinate.MoveLeft(1)
		case constants.East:
			currentCoordinate = currentCoordinate.MoveRight(1)
		}
	}
	return int64(len(coordinatesVisited))
}
