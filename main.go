package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\033[1;1H\033[2J\033[91mDNA\033[0m \033[94mReconstruction\033[0m \033[33mv00.03\033[0m\n")
	// Nucleotide Length for Staple Partitions
	NUCLEOTIDE_LENGTH := 6

	// Get Staple Strands from staples.txt
	for _, s := range ReadLines("sample/staples.txt") {
		fmt.Printf("%s\n", s)
	}

	fmt.Printf("%c\n", MakeNucleotide('A').base)
	fmt.Printf("%s\n", MakeStrand("ACTG CAGT").Bases()[:NUCLEOTIDE_LENGTH])
	strand := MakeStrand("ACTG CAGT")
	fmt.Printf("%s\n", strand.Bases())
	fmt.Printf("%s\n", strand.Complement().Bases())
	fmt.Printf("%s\n", strand.Bases())

	scaffold := Scaffold{
		MakeStrand("AGTCGTCATGCA"),
		[][]Strand{
			[]Strand{
				MakeStrand("TCAG"),
				MakeStrand("CAGT"),
			},
		},
		[][]int{
			[]int{
				0,
				4,
			},
		},
	}
	fmt.Printf("%s\n", scaffold.scaffold.Bases())
	for sidx, staple := range scaffold.staples {
		for idx, strand := range staple {
			fmt.Printf("%v %s\n", scaffold.bonds[sidx][idx], strand.Bases())
		}
	}
	fmt.Printf("%s", scaffold.MatchString())
	scaffold = scaffold.BondStaple([]Strand{MakeStrand("ACGT"),}, []int{8,})
	fmt.Printf("%s", scaffold.MatchString())
	fmt.Printf("%v\n", scaffold.MatchStrand(MakeStrand("GT")))
}
