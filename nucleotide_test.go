package main

import (
	"fmt"
	"testing"
)

func TestNucleotideComplement(t *testing.T) {
	var tests = []struct {
		n Nucleotide
		want Nucleotide
	}{
		{MakeNucleotide('A'), MakeNucleotide('T')},
		{MakeNucleotide('C'), MakeNucleotide('G')},
		{MakeNucleotide('T'), MakeNucleotide('A')},
		{MakeNucleotide('G'), MakeNucleotide('C')},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%c", tt.n.base)
		t.Run(testname, func(t *testing.T) {
			ans := tt.n.Complement()
			if ans != tt.want {
				t.Errorf("got `%c`, want `%c`", ans.base, tt.want.base)
			}
		})
	}
}
