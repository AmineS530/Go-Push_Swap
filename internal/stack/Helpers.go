package stack

type Node struct {
	Val   int
	Index int
	next  *Node
	prev  *Node
}

type Stack struct {
	Head *Node
	Size int
}

func New() *Stack {
	return &Stack{}
}

func (n *Node) Next() *Node { return n.next }
func (n *Node) Prev() *Node { return n.prev }
