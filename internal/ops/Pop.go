package ops

import "push_swap/internal/stack"

// sa - swap top two of a
func Sa(a *stack.Stack) string {
	if a.Size >= 2 {
		a.Head.Val, a.Head.Next().Val = a.Head.Next().Val, a.Head.Val
		a.Head.Index, a.Head.Next().Index = a.Head.Next().Index, a.Head.Index
	}
	return "sa"
}

// sb - swap top two of b
func Sb(b *stack.Stack) string {
	if b.Size >= 2 {
		b.Head.Val, b.Head.Next().Val = b.Head.Next().Val, b.Head.Val
		b.Head.Index, b.Head.Next().Index = b.Head.Next().Index, b.Head.Index
	}
	return "sb"
}

func Ss(a, b *stack.Stack) string {
	Sa(a)
	Sb(b)
	return "ss"
}
