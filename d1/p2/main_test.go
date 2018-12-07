package main

import (
	"testing"
)

func TestFindRepatedFrequency(t *testing.T) {
	freqs := []struct {
		deltas []string
		sum int64
	}{
		{[]string{"+1","-1"}, 0},
		{[]string{"+3", "+3", "+4", "-2", "-4"}, 10},
		{[]string{"-6", "+3", "+8", "+5", "-6"}, 5},
		{[]string{"+7", "+7", "-2", "-7", "-4"}, 14},
	}

	for _, freq := range freqs {
		result := findRepeatedFrequency(freq.deltas)
		if result != freq.sum {
			t.Errorf("findRepeatedFrequency(%s) was incorrect, got: %d, expected: %d", freq.deltas, result, freq.sum)
		}
	}
}
