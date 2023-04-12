package join_test

import (
	"fmt"
	"github.com/krukkrz/generics/join"
	"testing"
)

func TestJoin(t *testing.T) {
	testCases := []struct {
		name      string
		elements  []person
		separator string
		expected  string
	}{
		{
			name: "join elements",
			elements: []person{
				{7, "Marry"},
				{9, "Harry"},
				{2, "Barry"},
				{10, "Garry"},
				{17, "Perry"},
			},
			separator: ",",
			expected:  "{age: 7, name: Marry},{age: 9, name: Harry},{age: 2, name: Barry},{age: 10, name: Garry},{age: 17, name: Perry}",
		},
		{
			name:      "0 elements",
			elements:  []person{},
			separator: ",",
			expected:  "",
		},
		{
			name: "one element in slice",
			elements: []person{
				{7, "Marry"},
			},
			separator: ",",
			expected:  "{age: 7, name: Marry}",
		},
		{
			name: "empty separator",
			elements: []person{
				{7, "Marry"},
				{9, "Harry"},
				{2, "Barry"},
				{10, "Garry"},
				{17, "Perry"},
			},
			separator: "",
			expected:  "{age: 7, name: Marry}{age: 9, name: Harry}{age: 2, name: Barry}{age: 10, name: Garry}{age: 17, name: Perry}",
		},
		{
			name: "special characters",
			elements: []person{
				{7, "Krzyś"},
				{9, "Harry"},
				{2, "Barry"},
				{10, "Garry"},
				{17, "Perry"},
			},
			separator: ",",
			expected:  "{age: 7, name: Krzyś},{age: 9, name: Harry},{age: 2, name: Barry},{age: 10, name: Garry},{age: 17, name: Perry}",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := join.Join(tc.elements, tc.separator)
			if actual != tc.expected {
				t.Errorf("extected: %s\ngot: %s\n", tc.expected, actual)
			}
		})
	}
}

type person struct {
	age  int
	name string
}

func (p person) String() string {
	return fmt.Sprintf("{age: %d, name: %s}", p.age, p.name)
}
