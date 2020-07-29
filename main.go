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
		fmt.Printf("%s\n", staple_strands[len(staple_strands)-1].Bases())
	}

	// Get scaffold length
	for _, s := range staple_strands {
		scaffold_length += s.Length()
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
	for j := 0; j < len(staple_strands); j++ {
		for i := 0; i < staple_strands[j].Length()/NUCLEOTIDE_LENGTH; i++ {
			SubStrand := MakeStrand(staple_strands[j].Bases()[(i) * NUCLEOTIDE_LENGTH:(i+1) * NUCLEOTIDE_LENGTH])
			fmt.Printf("\n%s\n", SubStrand.Bases())
			fmt.Printf("%v\n", m13mp18.MatchStrand(SubStrand))
			if len(m13mp18.MatchStrand(SubStrand)) > 0 {
				// Only use the first match of all of the matches
				matchPosition := m13mp18.MatchStrand(SubStrand)[0]
				m13mp18 = m13mp18.BondStaple([]Strand{SubStrand}, []int{matchPosition,})
			} else {
				m13mp18 = m13mp18.BondStaple([]Strand{SubStrand}, []int{-1,})
			}
			fmt.Printf("%s", m13mp18.MatchString())
		}
		if staple_strands[j].Length() % NUCLEOTIDE_LENGTH != 0 {
			SubStrand := MakeStrand(staple_strands[j].Bases()[int(staple_strands[j].Length()/NUCLEOTIDE_LENGTH)*NUCLEOTIDE_LENGTH:])
			fmt.Printf("\n%s\n", SubStrand.Bases())
			fmt.Printf("%v\n", m13mp18.MatchStrand(SubStrand))
			if len(m13mp18.MatchStrand(SubStrand)) > 0 {
				matchPosition := m13mp18.MatchStrand(SubStrand)[0]
				m13mp18 = m13mp18.BondStaple([]Strand{SubStrand}, []int{matchPosition,})
			}
			fmt.Printf("%s", m13mp18.MatchString())
		}
	}
	fmt.Printf("%v\n", m13mp18)
}
