package stack

type Node struct {
	val  int
	next *Node
	prev *Node
}

type Stack struct {
	head *Node
	size int
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Len() int {
	return s.size
}
