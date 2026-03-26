package stack

// ra - rotate a 
func Ra(a *Stack) string {
	if a.size >= 2 {
		a.head = a.head.next
	}
	return "ra"
}

// rb - rotate b
func Rb(b *Stack) string {
	if b.size >= 2 {
		b.head = b.head.next
	}
	return "rb"
}

func Rr(a, b *Stack) string {
	Ra(a)
	Rb(b)
	return "rr"
}
