package main

import (
	"fmt"
	"testing"
)

func CompareStrand(s1, s2 Strand) bool {
	if len(s1.bases) == len(s2.bases) {
		for i := 0; i < s1.Length(); i++ {
			if s1.bases[i] != s2.bases[i] {
				return false
			}
		}
		return true
	}
	return false
}

func TestStrandComplement(t *testing.T) {
	var tests = []struct {
		s Strand
		want Strand
	}{
		{MakeStrand("ACTG"), MakeStrand("TGAC")},
		{MakeStrand("TCGA"), MakeStrand("AGCT")},
		{MakeStrand("T G G A C"), MakeStrand("ACCTG")},
		{MakeStrand(""), MakeStrand("")},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.s.Bases())
		t.Run(testname, func(t *testing.T) {
			ans := tt.s.Complement()
			if !CompareStrand(ans, tt.want) {
				t.Errorf("got `%s`, want `%s`", ans.Bases(), tt.want.Bases())
			}
		})
	}
}

func TestBases(t *testing.T) {
	var tests = []struct {
		s Strand
		want string
	}{
		{MakeStrand("ACTG"), "ACTG"},
		{MakeStrand("TCGA"), "TCGA"},
		{MakeStrand("T G G A C"), "TGGAC"},
		{MakeStrand(""), ""},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.s.Bases())
		t.Run(testname, func(t *testing.T) {
			ans := tt.s.Bases()
			if ans != tt.want {
				t.Errorf("got `%s`, want `%s`", ans, tt.want)
			}
		})
	}
}

func TestLength(t *testing.T) {
	var tests = []struct {
		s Strand
		want int
	}{
		{MakeStrand("ACTG"), 4},
		{MakeStrand("TCGA"), 4},
		{MakeStrand("T G G A C"), 5},
		{MakeStrand("T"), 1},
		{MakeStrand(""), 0},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.s.Bases())
		t.Run(testname, func(t *testing.T) {
			ans := tt.s.Length()
			if ans != tt.want {
				t.Errorf("got `%d`, want `%d`", ans, tt.want)
			}
		})
	}
}
