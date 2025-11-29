package not_quite_lisp

type NotQuiteLisp interface {
	Part1(data string) int64
	Part2(data string) int64
}

type notQuiteLisp struct{}

func New() NotQuiteLisp {
	return &notQuiteLisp{}
}

func (n *notQuiteLisp) Part1(data string) int64 {
	return 0
}

func (n *notQuiteLisp) Part2(data string) int64 {
	return 0
}
