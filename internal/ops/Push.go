package ops

import "push_swap/internal/stack"

// pa - push top of b to a
func Pa(a, b *stack.Stack) string {
	val, ok := b.Pop()
	if ok {
		a.Push(val)
	}
	return "pa"
}

// pb - push top of a to b
func Pb(a, b *stack.Stack) string {
	val, ok := a.Pop()
	if ok {
		b.Push(val)
	}
	return "pb"
}
