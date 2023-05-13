package queue

import (
	"container/list"
)

type linkedListQueue struct {
	data *list.List
}

func newLinkedlistQueue() *linkedListQueue {
	return &linkedListQueue{
		data: list.New(),
	}
}

func (s *linkedListQueue) push(value any) {
	s.data.PushBack(value)
}

func (s *linkedListQueue) pop() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Front()
	s.data.Remove(e)
	return e.Value
}

func (s *linkedListQueue) peek() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Front()
	return e.Value
}

func (s *linkedListQueue) size() int {
	return s.data.Len()
}

func (s *linkedListQueue) isEmpty() bool {
	return s.data.Len() == 0
}

func (s *linkedListQueue) toList() *list.List {
	return s.data
}
