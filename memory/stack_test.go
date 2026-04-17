package memory

import "testing"

func TestStack_PushPop(t *testing.T) {
	s := NewStack()
	s.Push(0x0200)
	s.Push(0x0300)

	if got := s.Pop(); got != 0x0300 {
		t.Errorf("Pop: want 0x0300, got 0x%04X", got)
	}
	if got := s.Pop(); got != 0x0200 {
		t.Errorf("Pop: want 0x0200, got 0x%04X", got)
	}
}
