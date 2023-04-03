package linkedlist

import (
	"fmt"
	"testing"
)

type TestObject struct {
	TestField string
}

func TestInsert(t *testing.T) {
	to := generateTestObjects()
	ll := New[TestObject]()

	for _, obj := range to {
		ll.Insert(obj)
	}

	if ll.head == nil {
		t.Errorf("head should not be nil\n")
	}

	if ll.head.Next == nil {
		t.Errorf("head should have next value\n")
	}

	counter := 0
	currentNode := ll.head

	for {
		if currentNode.Value != to[len(to)-counter-1] {
			t.Errorf("unexpected value, got: %v, expected %v", currentNode.Value, to[counter])
		}

		if currentNode.Next != nil {
			currentNode = currentNode.Next
			counter++
			continue
		}

		if currentNode.Next == nil {
			break
		}
	}

	if counter != len(to)-1 {
		t.Errorf("incorrect number of elements, got: %d, expected: %d", counter, len(to))
	}

	if ll.length != len(to) {
		t.Errorf("incorrect list length, got: %d, expected: %d", counter, len(to))
	}
}

func TestGet(t *testing.T) {
}

func TestDelete(t *testing.T) {

}

func generateTestObjects() []TestObject {
	var to []TestObject
	for i := 0; i < 100; i++ {
		tf := fmt.Sprintf("object-%d", i)
		to = append(to, TestObject{tf})
	}
	return to
}

func (t TestObject) String() string {
	return fmt.Sprintf("TestObject{%s}", t.TestField)
}
