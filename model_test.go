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
