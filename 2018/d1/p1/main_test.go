package main

import (
	"testing"
)

func TestComputeFrequency(t *testing.T) {
	freqs := []struct {
		deltas []string
		sum int64
	}{
		{[]string{"+1","+1","+1"}, 3},
		{[]string{"+1", "+1", "-2"}, 0},
		{[]string{"-1", "-2", "-3"}, -6},
	}

	for _, freq := range freqs {
		result := computeFrequency(freq.deltas)
		if result != freq.sum {
			t.Errorf("computeFrequency(%s) was incorrect, got: %d, expected: %d", freq.deltas, result, freq.sum)
		}
	}
}
