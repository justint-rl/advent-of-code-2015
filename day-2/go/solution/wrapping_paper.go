package wrapping_paper

import (
	"fmt"

	dimension "github.com/justint-rl/advent-of-code-2015/day-2/go/model/dimension"
)

type WrappingPaper interface {
	Part1(data []string) int64
	Part2(data []string) int64
}

type wrappingPaperImpl struct{}

func New() WrappingPaper {
	return &wrappingPaperImpl{}
}

func (w *wrappingPaperImpl) Part1(data []string) int64 {
	totalWrappingPaper := int64(0)
	for _, d := range data {
		dim, err := dimension.FromString(d)
		if err != nil {
			fmt.Println("Error converting string to Dimension:", err)
		}
		totalWrappingPaper += (dim.SurfaceArea() + dim.SmallestSideArea())
	}
	return totalWrappingPaper
}

func (w *wrappingPaperImpl) Part2(data []string) int64 {
	return 0
}
