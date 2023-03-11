package main

import (
	"fmt"
	"testing"
	tetris "tetris/lib"
)

type Test struct {
	input    string
	expected int
}

func TestGoodInput(t *testing.T) {
	var tests = []Test{
		{
			input:    "input0",
			expected: 0,
		},
		{
			input:    "input1",
			expected: 9,
		},
		{
			input:    "input2",
			expected: 4,
		},
		{
			input:    "input3",
			expected: 5,
		},
		{
			input:    "input4",
			expected: 9,
		},
	}
	for _, test := range tests {
		content := tetris.ReadFileContent("./examples/" + test.input)
		square := tetris.Solve(content)
		output := GetNumberOfSpace(square)
		if output != test.expected {
			t.Error("❌ Test Failed:\nInputted: ", test.input, "\nExpected: ", test.expected, "\nReceived: ", output)
		} else {
			fmt.Println("✅ Test Failed:\nInputter: ", test.input, "\nExpected: ", test.expected)
		}
	}
}

func _TestBadInput(t *testing.T) {
	var tests = []Test{
		{
			input:    "badformat",
			expected: 0,
		},
		{
			input:    "badinput0",
			expected: 0,
		},
		{
			input:    "badinput1",
			expected: 0,
		},
		{
			input:    "badinput2",
			expected: 0,
		},
		{
			input:    "badinput3",
			expected: 0,
		},
		{
			input:    "badinput4",
			expected: 0,
		},
	}
	for _, test := range tests {
		content := tetris.ReadFileContent("./examples/" + test.input)
		square := tetris.Solve(content)
		output := GetNumberOfSpace(square)
		if output != test.expected {
			t.Error("Test Failed:\nInputted: ", test.input, "\nExpected: ", test.expected, "\nReceived: ", output)
		}
	}
}

func GetNumberOfSpace(grid [][]rune) int {
	count := 0
	for _, line := range grid {
		for _, block := range line {
			if block == '.' {
				count++
			}
		}
	}
	return count
}
