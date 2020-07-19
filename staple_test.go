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
		{MakeStaple(MakeStrand("ACTCTGAC"), 3), 2},
		{MakeStaple(MakeStrand("ACTCTG"), 3), 2},
		{MakeStaple(MakeStrand("ACTCTGACTG"), 3), 3},
		{MakeStaple(MakeStrand("A"), 3), 0},
		{MakeStaple(MakeStrand(""), 3), 0},
	}

	for idx, tt := range tests {
		testname := fmt.Sprintf("%v", idx)
		t.Run(testname, func(t *testing.T) {
			ans := tt.staple.Length()
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestStapleTotalLength(t *testing.T) {
	var tests = []struct {
		staple Staple
		want int
	}{
		{MakeStaple(MakeStrand("ACTCTGAC"), 3), 6},
		{MakeStaple(MakeStrand("ACTCTG"), 3), 6},
		{MakeStaple(MakeStrand("ACTCTGACTG"), 3), 9},
		{MakeStaple(MakeStrand("A"), 3), 0},
		{MakeStaple(MakeStrand(""), 3), 0},
	}

	for idx, tt := range tests {
		testname := fmt.Sprintf("%v", idx)
		t.Run(testname, func(t *testing.T) {
			ans := tt.staple.TotalLength()
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
