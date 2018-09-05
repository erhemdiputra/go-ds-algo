package stack

import (
	"container/list"
)

type Stack struct {
	list *list.List
}

func New() *Stack {
	s := Stack{
		list: list.New(),
	}
	return &s
}

func (s *Stack) Push(item interface{}) {
	s.list.PushFront(item)
}

func (s *Stack) Pop() interface{} {
	if s.Len() > 0 {
		e := s.list.Front()
		s.list.Remove(e)
		return e.Value
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	if s.Len() > 0 {
		return s.list.Front().Value
	}
	return nil
}

func (s *Stack) Len() int {
	return s.list.Len()
}
