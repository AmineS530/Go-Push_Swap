package solver

import (
	"push_swap/internal/stack"
)

func IsSorted(a *stack.Stack) bool {
	if a.Size <= 1 {
		return true
	}
	curr := a.Head
	for i := 0; i < a.Size-1; i++ {
		if curr.Index > curr.Next().Index {
			return false
		}
		curr = curr.Next()
	}
	return true
}
