package main

import (
	"fmt"
)

func ReverseString(str string) string {
	outString := ""
	for _, s := range str {
		outString = string(s) + outString
	}
	return outString
}

func main() {
	fmt.Printf("\033[1;1H\033[2J\033[91mDNA\033[0m \033[94mReconstruction\033[0m \033[33mv00.03\033[0m\n")
	// Nucleotide Length for Staple Partitions
	NUCLEOTIDE_LENGTH := 8
	var staple_strands []Strand
	scaffold_length := 0

	// Get Staple Strands from staples.txt
	for _, s := range ReadLines("sample/staples.txt") {
		staple_strands = append(staple_strands, MakeStrand(s))
		scaffold_length += len(s)
		fmt.Printf("%s\n", staple_strands[len(staple_strands)-1].Bases())
	}

	// Make sure that we don't go beyond the end of the strand
	if scaffold_length > len(m13mp18f.Bases()) {
		scaffold_length = len(m13mp18f.Bases())
	}

	//TGATAGACGGTTTTTCGCCC
	m13mp18 := Scaffold{
		MakeStrand(ReverseString(m13mp18f.Bases()[:scaffold_length])),
		[][]Strand{ },
		[][]int{ },
	}

	// Match substrands from staple strands
	for i := 0; i < staple_strands[0].Length()/NUCLEOTIDE_LENGTH; i++ {
		SubStrand := MakeStrand(staple_strands[0].Bases()[(i) * NUCLEOTIDE_LENGTH:(i+1) * NUCLEOTIDE_LENGTH])
		fmt.Printf("\n%s\n", SubStrand.Bases())
		fmt.Printf("%v\n", m13mp18.MatchStrand(SubStrand))
		matchPosition := m13mp18.MatchStrand(SubStrand)[0]
		m13mp18 = m13mp18.BondStaple([]Strand{SubStrand}, []int{matchPosition,})
		fmt.Printf("%s", m13mp18.MatchString())
	}
}
