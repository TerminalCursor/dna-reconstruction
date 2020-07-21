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
		{Scaffold{MakeStrand("ACTGTGAC"),[][]Strand{[]Strand{MakeStrand("TG"),},},[][]int{[]int{6,}}}, MakeStrand("TG"), []int{0,}},
		{Scaffold{MakeStrand("ACTGTGAC"),[][]Strand{[]Strand{MakeStrand("TG"),},},[][]int{[]int{0,}}}, MakeStrand("CTG"), []int{5,}},
		{Scaffold{MakeStrand("ACTGTGAC"),[][]Strand{[]Strand{MakeStrand("TG"),},},[][]int{[]int{-1,}}}, MakeStrand("TG"), []int{0, 6,}},
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

func TestInIntList(t *testing.T) {
	var tests = []struct {
		val int
		list []int
		want bool
	}{
		{0, []int{0, 1, 4, 3, 5}, true},
		{1, []int{0, 1, 4, 3, 5}, true},
		{2, []int{0, 1, 4, 3, 5}, false},
		{3, []int{0, 1, 4, 3, 5}, true},
		{4, []int{0, 1, 4, 3, 5}, true},
		{5, []int{0, 1, 4, 3, 5}, true},
		{6, []int{0, 1, 4, 3, 5}, false},
		{7, []int{0, 1, 4, 3, 5}, false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.val)
		t.Run(testname, func(t *testing.T) {
			ans := InIntList(tt.val, tt.list)
			if ans != tt.want {
				t.Errorf("got `%t`, want `%t`", ans, tt.want)
			}
		})
	}
}
