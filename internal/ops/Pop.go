package ops

import "push_swap/internal/stack"

// sa - swap top two of a
func Sa(a *stack.Stack) string {
	if a.Size >= 2 {
		a.Head.Val, a.Head.Next().Val = a.Head.Next().Val, a.Head.Val
	}
	return "sa"
}

// sb - swap top two of b
func Sb(b *stack.Stack) string {
	if b.Size >= 2 {
		b.Head.Val, b.Head.Next().Val = b.Head.Next().Val, b.Head.Val
	}
	return "sb"
}

func Ss(a, b *stack.Stack) string {
	Sa(a)
	Sb(b)
	return "ss"
}
