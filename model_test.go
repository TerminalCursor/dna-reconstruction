package main

import (
	"fmt"
	"testing"
)

func TestComplement(t *testing.T) {
	var tests = []struct {
		c Nucleotide
		want Nucleotide
	}{
		{Nucleotide{0x41}, Nucleotide{0x54}},
		{Nucleotide{0x54}, Nucleotide{0x41}},
		{Nucleotide{0x43}, Nucleotide{0x47}},
		{Nucleotide{0x47}, Nucleotide{0x43}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.c.base)
		t.Run(testname, func(t *testing.T) {
			ans := tt.c.Complement()
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans.base, tt.want.base)
			}
		})
	}
}

func TestStrandLength(t *testing.T) {
	var tests = []struct {
		bases string
		want int
	}{
		{"ATCG", 4},
		{"A TCG", 4},
		{"TAGC", 4},
		{"A T C G", 4},
		{"A", 1},
		{"", 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.bases)
		t.Run(testname, func(t *testing.T) {
			ans := MakeStrand(tt.bases).Length()
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func CompareSlice(i, j []int) bool {
	if len(i) == len(j) {
		for k := 0; k < len(i); k++ {
			if i[k] != j[k] {
				return false
			}
		}
		return true
	}
	return false
}

func TestMatch(t *testing.T) {
	thescaffold := MakeStrand("TGAT AGAC GGTT TTTC GCCC TTTG ACGT TGGA GTCC ACGT TCTT TAAT AGTG GACT CTTG")
	var tests = []struct {
		staple Strand
		scaffold Strand
		want []int
	}{
		{MakeStrand("ACTA"), thescaffold, []int{-1, 0}},
		{MakeStrand("AAC"), thescaffold, []int{-1, 21, 27, 57}},
		{MakeStrand("CCAA"), thescaffold, []int{-1, 8}},
		{MakeStrand(""), MakeStrand(""), []int{-1}},
		{MakeStrand("A"), MakeStrand(""), []int{-1}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.staple.Bases())
		t.Run(testname, func(t *testing.T) {
			ans := tt.staple.Match(tt.scaffold)
			if !CompareSlice(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
