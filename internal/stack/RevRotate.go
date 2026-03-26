package stack

// rra - reverse rotate a (bottom goes to top)
func Rra(a *Stack) string {
	if a.size >= 2 {
		a.head = a.head.prev // just move head backward, O(1)
	}
	return "rra"
}

// rrb - reverse rotate b
func Rrb(b *Stack) string {
	if b.size >= 2 {
		b.head = b.head.prev
	}
	return "rrb"
}

func Rrr(a, b *Stack) string {
	Rra(a)
	Rrb(b)
	return "rrr"
}
