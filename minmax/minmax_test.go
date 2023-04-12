package minmax_test

import (
	"github.com/krukkrz/generics/minmax"
	"testing"
)

func TestMin_String(t *testing.T) {
	max := "book"
	min := "apple"
	result := minmax.Min(min, max)
	if result != min {
		t.Errorf("expecting: %v, got: %v", min, result)
	}
}

func TestMin_StringsEqual(t *testing.T) {
	max := "book"
	min := "book"
	result := minmax.Min(min, max)
	if result != min {
		t.Errorf("expecting: %v, got: %v", min, result)
	}
}

func TestMin_Int(t *testing.T) {
	max := 66
	min := 21
	result := minmax.Min(min, max)
	if result != min {
		t.Errorf("expecting: %v, got: %v", min, result)
	}
}

func TestMin_IntsEqual(t *testing.T) {
	max := 66
	min := 66
	result := minmax.Min(min, max)
	if result != min {
		t.Errorf("expecting: %v, got: %v", min, result)
	}
}

func TestMax_StringsEqual(t *testing.T) {
	max := "book"
	min := "book"
	result := minmax.Max(min, max)
	if result != max {
		t.Errorf("expecting: %v, got: %v", max, result)
	}
}

func TestMax_Int(t *testing.T) {
	max := 66
	min := 21
	result := minmax.Max(min, max)
	if result != max {
		t.Errorf("expecting: %v, got: %v", max, result)
	}
}

func TestMax_IntsEqual(t *testing.T) {
	max := 66
	min := 66
	result := minmax.Max(min, max)
	if result != max {
		t.Errorf("expecting: %v, got: %v", max, result)
	}
}
