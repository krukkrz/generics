package linkedlist

type Interface[T any] interface {
	Get(index int) T
	Insert(T)
	Delete(index int)
}

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	head   *Node[T]
	length int
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head:   nil,
		length: 0,
	}
}

func (l LinkedList[T]) Get(index int) *T {
	if index < 0 {
		return nil
	}
	current := l.head
	for i := 0; i < index; i++ {
		next := current.Next
		if next == nil {
			return nil
		}
		current = next
	}
	return &current.Value
}

func (l *LinkedList[T]) Insert(t T) {
	n := &Node[T]{
		Value: t,
		Next:  l.head,
	}
	l.head = n
	l.length++
}

func (l *LinkedList[T]) Delete(index int) {
	if index == 0 {
		l.head = l.head.Next
		return
	}

	var prev *Node[T]
	var del *Node[T]

	current := l.head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	prev = current
	del = current.Next

	prev.Next = del.Next
}
