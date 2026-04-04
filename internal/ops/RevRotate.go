package ops

import "push_swap/internal/stack"

// rra - reverse rotate a (bottom goes to top)
func Rra(a *stack.Stack) string {
	if a.Size >= 2 {
		a.Head = a.Head.Prev() // just move head backward, O(1)
	}
	return "rra"
}

// rrb - reverse rotate b
func Rrb(b *stack.Stack) string {
	if b.Size >= 2 {
		b.Head = b.Head.Prev()
	}
	return "rrb"
}

func Rrr(a, b *stack.Stack) string {
	Rra(a)
	Rrb(b)
	return "rrr"
}
