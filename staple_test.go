package main

import (
	"fmt"
	"testing"
)

func TestStapleLength(t *testing.T) {
	var tests = []struct {
		staple Staple
		want int
	}{
		{MakeStaple(MakeStrand("ACTCTGAC"), 3), 3},
		{MakeStaple(MakeStrand("ACTCTG"), 3), 2},
		{MakeStaple(MakeStrand("ACTCTGACTG"), 3), 4},
		{MakeStaple(MakeStrand("A"), 3), 1},
		{MakeStaple(MakeStrand(""), 3), 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.staple.pieces)
		t.Run(testname, func(t *testing.T) {
			ans := tt.staple.Length()
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
