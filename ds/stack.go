package ds

type Stack struct {
    list LinkedList
}

func (s *Stack) Push(value string) {
	newNode := &Node{data: value, Next: nil}
	curHead := s.list.Head

	s.list.Head = newNode
	newNode.Next = curHead
}

func (s *Stack) Pop() (string, bool) {

	if s.list.Head == nil {
		return "", false
	}

	val := s.list.Head.data
	s.list.RemoveAt(0)

	return val, true
}