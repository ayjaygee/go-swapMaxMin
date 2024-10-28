package main

import (
	"slices"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected []int
	}{
		{"213", 213, []int{2, 1, 3}},
		{"12345", 12345, []int{1, 2, 3, 4, 5}},
		{"100", 100, []int{1, 0, 0}},
	}

	for _, test := range tests {
		t.Run(test.name, func(u *testing.T) {
			got := SplitInt(test.input)
			if slices.Compare(test.expected, got) != 0 {
				u.Errorf("Expected %v but got %v", test.expected, got)
			}
		})
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"213", []int{2, 1, 3}, 213},
		{"12345", []int{1, 2, 3, 4, 5}, 12345},
		{"100", []int{1, 0, 0}, 100},
	}

	for _, test := range tests {
		t.Run(test.name, func(u *testing.T) {
			got := JoinInt(test.input)
			if got != test.expected {
				u.Errorf("Expected %v but got %v", test.expected, got)
			}
		})
	}
}

func TestSwapElement(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		i, j     int
		expected []int
	}{
		{"213", []int{2, 1, 3}, 0, 1, []int{1, 2, 3}},
		{"12345", []int{1, 2, 3, 4, 5}, 2, 4, []int{1, 2, 5, 4, 3}},
		{"100", []int{1, 0, 0}, 1, 0, []int{0, 1, 0}},
	}

	for _, test := range tests {
		t.Run(test.name, func(u *testing.T) {
			got := SwapElements(test.input, test.i, test.j)
			if !slices.Equal(test.expected, got) {
				u.Errorf("Expected %v but got %v", test.expected, got)
			}
		})
	}
}

func TestSwapMaxMinForMax(t *testing.T) {
	tests := []struct {
		name  string
		input int
		exMax int
	}{
		{"213 - 3 digits no zeroes", 213, 312},
		{"12345 - 5 digits no zeroes", 12345, 52341},
		{"100 - 3 digits with zeroes", 100, 100},
	}

	for _, test := range tests {
		t.Run(test.name, func(u *testing.T) {
			gotMax, _ := SwapMaxMin(test.input)
			if test.exMax != gotMax {
				u.Errorf("Expected %v but got %v", test.exMax, gotMax)
			}
		})
	}
}

func TestSwapMaxMinForMin(t *testing.T) {
	tests := []struct {
		name  string
		input int
		exMin int
	}{
		{"213 - 3 digits no zeroes", 213, 123},
		{"12345 - 5 digits no zeroes", 12345, 12345},
		{"100 - 3 digits with zeroes", 100, 100},
	}

	for _, test := range tests {
		t.Run(test.name, func(u *testing.T) {
			_, gotMin := SwapMaxMin(test.input)
			if test.exMin != gotMin {
				u.Errorf("Expected %v but got %v", test.exMin, gotMin)
			}
		})
	}
}
