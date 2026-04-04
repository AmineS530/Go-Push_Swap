package stack

import "sort"

func (s *Stack) Push(val int) {
	node := &Node{Val: val}

	if s.Head == nil {
		node.next = node
		node.prev = node
		s.Head = node
	} else {
		tail := s.Head.prev

		node.next = s.Head
		node.prev = tail

		tail.next = node
		s.Head.prev = node

		s.Head = node
	}
	s.Size++
}

func (s *Stack) Pop() (int, bool) {
	if s.Head == nil {
		return 0, false
	}

	val := s.Head.Val

	if s.Size == 1 {
		s.Head = nil
	} else {
		tail := s.Head.prev

		s.Head = s.Head.next
		s.Head.prev = tail
		tail.next = s.Head
	}

	s.Size--
	return val, true
}

func Normalize(nums []int) []int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	rankMap := make(map[int]int)
	for i, v := range sorted {
		rankMap[v] = i
	}

	indices := make([]int, len(nums))
	for i, v := range nums {
		indices[i] = rankMap[v]
	}
	return indices
}

func NewFromSlice(nums []int, indices []int) *Stack {
	s := New()
	for i := len(nums) - 1; i >= 0; i-- {
		s.pushNode(&Node{
			Val:   nums[i],
			Index: indices[i],
		})
	}
	return s
}

func (s *Stack) pushNode(node *Node) {
	if s.Head == nil {
		node.next = node
		node.prev = node
		s.Head = node
	} else {
		tail := s.Head.prev
		node.next = s.Head
		node.prev = tail
		tail.next = node
		s.Head.prev = node
		s.Head = node
	}
	s.Size++
}
