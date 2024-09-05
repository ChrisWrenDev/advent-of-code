package main

import (
	"testing"
)

func TestSolvePart1(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int64
	}{
		{"test 1", []string{"2x3x4"}, 58},
		{"test 2", []string{"1x1x10"}, 43},
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
		input    []string
		expected int64
	}{
		{"test 1", []string{"2x3x4"}, 34},
		{"test 2", []string{"1x1x10"}, 14},
	}

	for _, test := range tests {
		result := solvePart2(test.input)
		if result != test.expected {
			t.Errorf("solvePart2: expected %d but got %d", test.expected, result)
		}
	}
}
