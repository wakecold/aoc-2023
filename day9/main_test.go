package main

import "testing"

type TestCase struct {
	input    []int
	expected int
}

func TestPredict(t *testing.T) {
	testCases := []TestCase{
		{
			[]int{0, 3, 6, 9, 12, 15},
			18,
		},
		{
			[]int{1, 3, 6, 10, 15, 21},
			28,
		},
		{
			[]int{10, 13, 16, 21, 30, 45},
			68,
		},
	}

	for _, c := range testCases {
		result := predict(c.input)
		if result != c.expected {
			t.Fatalf("Incorrect result on case %v - expected: %v, got: %v", c.input, c.expected, result)
		}
	}

}

func TestPredictBackwards(t *testing.T) {
	arr := []int{10, 13, 16, 21, 30, 45}
	expected := 5
	result := predictBackwards(arr)
	if result != expected {
		t.Fatalf("Incorrect result on backwards predict: expected-%v,got-%v", expected, result)
	}
}
