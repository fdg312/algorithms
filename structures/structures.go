package structures

import "fmt"

type Node[T any] struct {
	Value    T
	previous *Node[T]
	next     *Node[T]
}

type Stacker[T any] interface {
	Peek() (T, error)
	Push(value T)
	Pop() (T, error)
	Size() int32
}

type Stack[T any] struct {
	head *Node[T]
	size int32
}

func (s *Stack[T]) Peek() (T, error) {
	if s.head == nil {
		var zero T
		return zero, fmt.Errorf("head is not set")
	}
	return s.head.Value, nil
}

func (s *Stack[T]) Push(value T) {
	s.size++

	s.head = &Node[T]{
		Value:    value,
		previous: s.head,
	}
}

func (s *Stack[T]) Pop() (T, error) {
	if s.head == nil {
		var zero T
		return zero, fmt.Errorf("you cannot pop from an empty stack")
	}

	s.size--
	previous := s.head
	s.head = s.head.previous
	return previous.Value, nil

}

func (s *Stack[T]) Size() int32 {
	return s.size
}

func NewStack[T any]() Stacker[T] {
	return &Stack[T]{
		head: nil,
		size: 0,
	}
}

type Queuer[T any] interface {
	Enqueue(T)
	Dequeue() (T, error)
	Size() int32
	Peek() (T, error)
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int32
}

func NewQueue[T any]() Queuer[T] {
	return &Queue[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (qs *Queue[T]) Enqueue(val T) {
	qs.size++
	node := &Node[T]{Value: val}

	if qs.head == nil {
		qs.tail = node
		qs.head = node
		return
	}

	qs.tail.next = node

	qs.tail = node

}

func (qs *Queue[T]) Dequeue() (T, error) {
	if qs.head == nil {
		var zero T
		return zero, fmt.Errorf("Queue is empty")
	}

	qs.size--
	curr := qs.head
	qs.head = qs.head.next

	if qs.head == nil {
		qs.tail = nil
	}

	return curr.Value, nil

}

func (qs *Queue[T]) Size() int32 {
	return qs.size
}

func (qs *Queue[T]) Peek() (T, error) {
	if qs.head == nil {
		var zero T
		return zero, fmt.Errorf("Queue is empty")
	}
	return qs.head.Value, nil
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
