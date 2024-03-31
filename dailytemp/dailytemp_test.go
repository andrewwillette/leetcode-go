package dailytemp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Peek() int {
	return s[len(s)-1]
}

func (s stack) IsEmpty() bool {
	return len(s) == 0
}

// https://leetcode.com/problems/daily-temperatures/
func dailyTemperatures(temperatures []int) []int {
	result := []int{}
	// use a stack to keep track of the index of the temperature
	s := new(stack)
	// populate memo
	for i, vi := range temperatures {
		// if the stack is empty, push the index
		if s.IsEmpty() {
			*s = s.Push(i)
			continue
		}
		fmt.Printf("stack : %v\nvi: %v\n\n", s, vi)
		for !s.IsEmpty() {
			peek := s.Peek()
			fmt.Printf("comparing %v < %v ?\n\n", temperatures[peek], vi)
			if temperatures[peek] < vi {
				*s, _ = s.Pop()
				result = append(result, i-peek)
				fmt.Printf("updating result: %v\n\n", result)
			} else {
				break
			}
		}
		*s = s.Push(i)
	}
	fmt.Printf("s: %v\n\n", s)
	return result
}

func TestDailyTemperatures(t *testing.T) {
	// test case 1
	// temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	// expected := []int{1, 1, 4, 2, 1, 1, 0, 0}
	// result := dailyTemperatures(temperatures)
	// require.Equal(t, expected, result)

	// // test case 2
	// temperatures = []int{30, 40, 50, 60}
	// expected = []int{1, 1, 1, 0}
	// result = dailyTemperatures(temperatures)
	// require.Equal(t, expected, result)
	// // test case 3

	//
	// (i = 0 , s = [0], peek = 0), (i = 1, s
	temperatures := []int{30, 60, 90}
	expected := []int{1, 1, 0}
	result := dailyTemperatures(temperatures)
	require.Equal(t, expected, result)
}
