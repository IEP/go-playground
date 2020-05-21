package main

import "testing"

type testSet struct {
	input  float64
	output int
}

func TestTrunc(t *testing.T) {
	test := []testSet{
		{12.345, 12},
		{13.111, 13},
	}
	for _, value := range test {
		truncated := trunc(value.input)
		if truncated != value.output {
			t.Errorf("trunc(%f) != %d", value.input, value.output)
		}
	}
}
