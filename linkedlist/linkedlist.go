package linkedlist

type LinkedListInterface[T any] interface {
	Get(index int) T
	Insert(T)
	Delete(index int)
	String()
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

func (l LinkedList[T]) Get(index int) T {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}

func (l LinkedList[T]) String() {
	//TODO implement me
	panic("implement me")
}
