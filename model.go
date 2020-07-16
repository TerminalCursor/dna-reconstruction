package main

import (
	"fmt"
)

type Nucleotide struct {
	base byte
}

type Strand struct {
	strand []Nucleotide
}

func (n Nucleotide) Complement() Nucleotide {
	var outN Nucleotide
	switch n.base {
		case 0x41:
			outN.base = 0x54
		case 0x54:
			outN.base = 0x41
		case 0x43:
			outN.base = 0x47
		case 0x47:
			outN.base = 0x43
		default:
			outN.base = 0x20
	}
	return outN
}

func (s Strand) Length() int {
	return len(s.strand)
}

func (n Nucleotide) Print() {
	fmt.Printf("%c", n.base)
}

func (s Strand) Print() {
	for i := 0; i < s.Length(); i++ {
		fmt.Printf("%c", s.strand[i].base)
	}
	fmt.Printf("\n")
}
