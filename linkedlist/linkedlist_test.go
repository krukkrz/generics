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
	to = revert(to)
	for {
		if currentNode.Value != to[counter] {
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
	testCases := []struct {
		name          string
		index         int
		expectedValue *TestObject
	}{
		{
			name: "iterate through all",
		},
		{
			name:          "index out of range",
			index:         200,
			expectedValue: nil,
		},
		{
			name:          "index negative",
			index:         -1,
			expectedValue: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			to := generateTestObjects()
			ll := New[TestObject]()

			for _, obj := range to {
				ll.Insert(obj)
			}
			to = revert(to)

			iterateThroughAll := tc.index == 0

			if iterateThroughAll {
				for i, o := range to {
					actual := ll.Get(i)
					if o != *actual {
						t.Errorf("unexpected value, got: %v, expected: %v", actual, o)
					}
				}
			} else {
				actual := ll.Get(tc.index)
				if actual != tc.expectedValue {
					t.Errorf("unexpected value, got: %v, expected: %v", actual, tc.expectedValue)
				}
			}
		})
	}

}

func TestDelete(t *testing.T) {
	testCases := []struct {
		name  string
		index int
	}{
		{"remove first element", 0},
		{"remove 20th element", 20},
		{"remove last element", 99},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			to := generateTestObjects()
			ll := New[TestObject]()

			for _, obj := range to {
				ll.Insert(obj)
			}
			to = revert(to)

			o := ll.Get(tc.index)
			ll.Delete(tc.index)

			for i, _ := range to {
				actual := ll.Get(i)

				if actual == o {
					t.Errorf("object with index %d was not deleted, and now has index: %d", tc.index, i)
				}
			}
		})
	}
}

func generateTestObjects() []TestObject {
	var to []TestObject
	for i := 0; i < 100; i++ {
		tf := fmt.Sprintf("object-%d", i)
		to = append(to, TestObject{tf})
	}
	return to
}

func revert(l []TestObject) []TestObject {
	var result []TestObject
	for i := len(l) - 1; i >= 0; i-- {
		result = append(result, l[i])
	}
	return result
}

func (t TestObject) String() string {
	return fmt.Sprintf("TestObject{%s}", t.TestField)
}
