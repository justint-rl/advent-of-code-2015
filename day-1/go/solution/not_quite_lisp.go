package not_quite_lisp

type NotQuiteLisp interface {
	Part1(data string) int64
	Part2(data string) int64
}

type notQuiteLisp struct{}

func New() NotQuiteLisp {
	return &notQuiteLisp{}
}

const UP_FLOOR rune = '('
const DOWN_FLOOR rune = ')'

func (n *notQuiteLisp) Part1(data string) int64 {
	counter := int64(0)
	for _, char := range data {
		switch char {
		case DOWN_FLOOR:
			counter -= 1
		case UP_FLOOR:
			counter += 1
		}
	}
	return counter
}

func (n *notQuiteLisp) Part2(data string) int64 {
	counter := int64(0)

	for idx, char := range data {
		switch char {
		case UP_FLOOR:
			counter += 1
		case DOWN_FLOOR:
			counter -= 1
		}
		if counter < 0 {
			return int64(idx + 1)
		}
	}

	return -1
}
