package solution

import (
	"strings"

	constants "github.com/justint-rl/advent-of-code-2015/day-3/go/constants"
	model "github.com/justint-rl/advent-of-code-2015/day-3/go/model"
)

type HouseVisits interface {
	Part1(route string) int64
	Part2(route string) int64
}

type houseVisitsImpl struct{}

func New() HouseVisits {
	return &houseVisitsImpl{}
}

func (h *houseVisitsImpl) GetHouseVisited(route string) map[model.Coordinate]struct{} {
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
	return coordinatesVisited
}

func (h *houseVisitsImpl) Part1(route string) int64 {
	coordinatesVisited := h.GetHouseVisited(route)
	return int64(len(coordinatesVisited))
}

func (h *houseVisitsImpl) Part2(route string) int64 {
	var santaRoute strings.Builder
	var roboSantaRoute strings.Builder

	for i, c := range route {
		if i%2 == 1 {
			santaRoute.WriteRune(c)
		} else {
			roboSantaRoute.WriteRune(c)
		}
	}
	santaHouseVisited := h.GetHouseVisited(santaRoute.String())
	roboSantaHouseVisisted := h.GetHouseVisited(roboSantaRoute.String())

	mergeHouseVisited := make(map[model.Coordinate]struct{})
	for k, v := range santaHouseVisited {
		mergeHouseVisited[k] = v
	}
	for k, v := range roboSantaHouseVisisted {
		mergeHouseVisited[k] = v
	}
	return int64(len(mergeHouseVisited))
}
