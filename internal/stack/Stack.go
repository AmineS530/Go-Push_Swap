package stack

func (s *Stack) Push(val int) {
	node := &Node{val: val}

	if s.head == nil {
		// first element points to itself in both directions
		node.next = node
		node.prev = node
		s.head = node
	} else {
		tail := s.head.prev // bottom of stack

		node.next = s.head // new node points forward to old head
		node.prev = tail   // new node points back to tail

		tail.next = node   // old tail points forward to new node
		s.head.prev = node // old head points back to new node

		s.head = node // new node is now the top
	}
	s.size++
}

func (s *Stack) Pop() (int, bool) {
	if s.head == nil {
		return 0, false
	}

	val := s.head.val

	if s.size == 1 {
		s.head = nil
	} else {
		tail := s.head.prev

		s.head = s.head.next // move head forward
		s.head.prev = tail   // new head points back to tail
		tail.next = s.head   // tail points forward to new head
	}

	s.size--
	return val, true
}
