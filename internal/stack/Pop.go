package stack

// sa - swap top two of a
func Sa(a *Stack) string {
	if a.size >= 2 {
		a.head.val, a.head.next.val = a.head.next.val, a.head.val
	}
	return "sa"
}

// sb - swap top two of b
func Sb(b *Stack) string {
	if b.size >= 2 {
		b.head.val, b.head.next.val = b.head.next.val, b.head.val
	}
	return "sb"
}

func Ss(a, b *Stack) string {
	Sa(a)
	Sb(b)
	return "ss"
}
