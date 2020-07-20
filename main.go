package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\033[91mDNA\033[0m \033[94mReconstruction\033[0m \033[33mv00.03\033[0m\n")
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
}
