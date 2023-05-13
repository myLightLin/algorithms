package queue

import (
	"container/list"
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := list.New()

	queue.PushBack(1)
	queue.PushBack(3)
	queue.PushBack(2)
	fmt.Print("队列 queue =")
	PrintList(queue)
}

func PrintList(list *list.List) {
	e := list.Front()
	fmt.Print("[")
	for e.Next() != nil {
		fmt.Print(e.Value, " ")
		e = e.Next()
	}
	fmt.Print(e.Value, "]\n")
}
