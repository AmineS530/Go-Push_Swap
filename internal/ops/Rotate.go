package ops

import "push_swap/internal/stack"

// ra - rotate a
func Ra(a *stack.Stack) string {
	if a.Size >= 2 {
		a.Head = a.Head.Next()
	}
	return "ra"
}

// rb - rotate b
func Rb(b *stack.Stack) string {
	if b.Size >= 2 {
		b.Head = b.Head.Next()
	}
	return "rb"
}

func Rr(a, b *stack.Stack) string {
	Ra(a)
	Rb(b)
	return "rr"
}
