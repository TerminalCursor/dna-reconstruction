package main

import (
	"fmt"
	"testing"
)

func CompareLists(l1, l2 []int) bool {
	if len(l1) == len(l2) {
		for i := 0; i < len(l1); i++ {
			if l1[i] != l2[i] {
				return false
			}
		}
		return true
	}
	return false
}

func TestMatchStrand(t *testing.T) {
	var tests = []struct {
		sc Scaffold
		s Strand
		want []int
	}{
		{Scaffold{MakeStrand("ACTGTGAC"),[][]Strand{},[][]int{}}, MakeStrand("TG"), []int{0, 6,}},
		{Scaffold{MakeStrand("ACTGTGAC"),[][]Strand{},[][]int{}}, MakeStrand("AC"), []int{2, 4,}},
		{Scaffold{MakeStrand("ACTGTGAC"),[][]Strand{},[][]int{}}, MakeStrand(""), []int{}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.s.Bases())
		t.Run(testname, func(t *testing.T) {
			ans := tt.sc.MatchStrand(tt.s)
			if !CompareLists(ans, tt.want) {
				t.Errorf("got `%v`, want `%v`", ans, tt.want)
			}
		})
	}
}
