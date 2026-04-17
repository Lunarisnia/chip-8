package memory

type Stack interface {
	Push(value uint16)
	Pop() uint16
}

type stack struct {
	container []uint16
}

func NewStack() Stack {
	return &stack{
		container: make([]uint16, 0),
	}
}

func (s *stack) Push(value uint16) {
	s.container = append(s.container, value)
}

func (s *stack) Pop() uint16 {
	lastValue := s.container[len(s.container)-1]
	s.container = s.container[:len(s.container)-1]
	return lastValue
}
