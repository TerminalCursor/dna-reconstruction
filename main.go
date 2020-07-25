package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\033[1;1H\033[2J\033[91mDNA\033[0m \033[94mReconstruction\033[0m \033[33mv00.03\033[0m\n")
	// Nucleotide Length for Staple Partitions
	//NUCLEOTIDE_LENGTH := 6
	var staple_strands []Strand
	scaffold_length := 0

	// Get Staple Strands from staples.txt
	for _, s := range ReadLines("sample/staples.txt") {
		staple_strands = append(staple_strands, MakeStrand(s))
		scaffold_length += len(s)
		fmt.Printf("%s\n", staple_strands[len(staple_strands)-1].Bases())
	}

	if scaffold_length > len(m13mp18f.Bases()) {
		scaffold_length = len(m13mp18f.Bases())
	}

	//TGATAGACGGTTTTTCGCCC
	m13mp18 := Scaffold{
		MakeStrand(m13mp18f.Bases()[:scaffold_length]),
		[][]Strand{
//			[]Strand{
//				MakeStrand("TCAG"),
//				MakeStrand("CAGT"),
//			},
		},
		[][]int{
//			[]int{
//				0,
//				-1,
//			},
		},
	}

	fmt.Printf("%s", m13mp18.MatchString())
	fmt.Printf("%v\n", m13mp18.MatchStrand(MakeStrand("ACTA")))
	m13mp18 = m13mp18.BondStaple([]Strand{MakeStrand("CCAA"),MakeStrand("AAAG")}, []int{8,12,})
	fmt.Printf("%s", m13mp18.MatchString())
	fmt.Printf("%v\n", m13mp18.MatchStrand(MakeStrand("GG")))
}
