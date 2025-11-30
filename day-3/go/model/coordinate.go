package model

type Coordinate struct {
	X int64
	Y int64
}

func New(x int64, y int64) Coordinate {
	return Coordinate{X: x, Y: y}
}

func (c Coordinate) MoveUp(steps int64) Coordinate {
	return New(c.X, c.Y+steps)
}

func (c Coordinate) MoveDown(steps int64) Coordinate {
	return New(c.X, c.Y-steps)
}

func (c Coordinate) MoveLeft(steps int64) Coordinate {
	return New(c.X-steps, c.Y)
}

func (c Coordinate) MoveRight(steps int64) Coordinate {
	return New(c.X+steps, c.Y)
}
