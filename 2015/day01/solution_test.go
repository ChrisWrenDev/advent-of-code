package main

import (
	"testing"
)

func TestSolvePart1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"test 1", "(())", 0},
		{"test 2", "(()(()(", 3},
		{"test 3", "())", -1},
		{"test 4", ")())())", -3},
	}

	for _, test := range tests {
		result := solvePart1(test.input)
		if result != test.expected {
			t.Errorf("solvePart1: expected %d but got %d", test.expected, result)
		}
	}
}

func TestSolvePart2(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"test 1", ")", 1},
		{"test 2", "()())", 5},
	}

	for _, test := range tests {
		result := solvePart2(test.input)
		if result != test.expected {
			t.Errorf("solvePart2: expected %d but got %d", test.expected, result)
		}
	}
}
