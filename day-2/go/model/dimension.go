package dimension

import (
	"fmt"
	"strconv"
	"strings"
)

type Dimension struct {
	Length int64
	Width  int64
	Height int64
}

const DIMENSION_DELIMITER = "x"

func FromString(s string) (Dimension, error) {
	dims := strings.Split(s, DIMENSION_DELIMITER)
	length, err := strconv.ParseInt(dims[0], 10, 64)
	if err != nil {
		fmt.Println("Error converting string to length:", err)
		return Dimension{}, err
	}
	width, err := strconv.ParseInt(dims[1], 10, 64)
	if err != nil {
		fmt.Println("Error converting string to width:", err)
		return Dimension{}, err
	}
	height, err := strconv.ParseInt(dims[2], 10, 64)
	if err != nil {
		fmt.Println("Error converting string to height:", err)
		return Dimension{}, err
	}
	return Dimension{Length: length, Width: width, Height: height}, nil
}

func (d Dimension) SurfaceArea() int64 {
	total := int64(0)
	total += 2 * (d.Length * d.Width)
	total += 2 * (d.Width * d.Height)
	total += 2 * (d.Height * d.Length)
	return total
}

func (d Dimension) SmallestSideArea() int64 {
	return min(
		(d.Length * d.Width),
		(d.Width * d.Height),
		(d.Height * d.Length),
	)
}

func (d Dimension) SmallestPerimeter() int64 {
	return min(
		(2*d.Length + 2*d.Width),
		(2*d.Width + 2*d.Height),
		(2*d.Height + 2*d.Length),
	)
}

func (d Dimension) Volume() int64 {
	return d.Length * d.Width * d.Height
}
