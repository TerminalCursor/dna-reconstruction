package main

import (
	"fmt"
	"os"
)

func PrintMatch(scaffold, staplePart Strand, offset int) {
	fmt.Printf("%s\n", scaffold.Bases())
	for i := 0; i < offset + staplePart.Length(); i++ {
		if i >= offset {
			fmt.Printf("%c", staplePart.Bases()[i-offset])
		} else {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
}

func GetPermutations(input [][]int, prep []int, WINDOW_SIZE int) [][]int {
	var out [][]int
	if(len(input) > 1) {
		if len(input[0]) > 0 {
			for i := 0; i < len(input[0]); i++ {
				validPath := true
				for j := 0; j < len(prep); j++ {
					if prep[j] <= input[0][i] && input[0][i] < (prep[j] + WINDOW_SIZE) && prep[j] != -1 {
						validPath = false
						break
					}
				}
				if validPath {
					perms := GetPermutations(input[1:], append(prep, input[0][i]), WINDOW_SIZE)
					for j := 0; j < len(perms); j++ {
						out = append(out, perms[j])
					}
				}
			}
		} else {
			fmt.Printf("EMPTY MATCH DETECTED\nTRIVIAL MATCH '-1' NOT INCLUDED\n")
			os.Exit(1)
		}
	} else {
		for i := 0; i < len(input[0]); i++ {
			validPath := true
			for j := 0; j < len(prep); j++ {
				if prep[j] <= input[0][i] && input[0][i] < (prep[j] + WINDOW_SIZE) && prep[j] != -1 {
					validPath = false
					break
				}
			}
			if validPath {
				out = append(out, append(prep, input[0][i]))
			}
		}
	}
	return out
}

/* https://en.wikipedia.org/wiki/ANSI_escape_code */

func main() {
	fmt.Printf("\033[91mDNA\033[0m \033[94mReconstruction\033[0m \033[33mv00.01\033[0m\n")
	//m13mp18 := MakeStrand("TGAT AGAC GGTT TTTC GCCC TTTG ACGT TGGA GTCC ACGT TCTT TAAT AGTG GACT CTTG")
	WINDOW_SIZE := 3
	eStaple := MakeStaple(MakeStrand("ACTATCTG"), WINDOW_SIZE)
	m13mp18 := MakeStrand(m13mp18f.Bases()[:eStaple.TotalLength()])
	strandConfs := make([][]int, eStaple.Length())
	for i := 0; i < eStaple.Length(); i++ {
		matches := eStaple.pieces[i].Match(m13mp18)
		strandConfs[i] = matches
		fmt.Printf("%s %v\n", eStaple.pieces[i].Bases(), matches)
	}
	fmt.Printf("%v\n", strandConfs)
	possibleCombinations := GetPermutations(strandConfs, []int{}, WINDOW_SIZE)
	fmt.Printf("%v\n", possibleCombinations)
	var mS []MatchedScaffold
	for i := 0; i < len(possibleCombinations); i++ {
		mS = append(mS, MatchedScaffold{m13mp18, eStaple, possibleCombinations[i]})
	}
	for i := 0; i < len(mS); i++ {
		fmt.Printf("%s", mS[i].MatchedString())
	}
}
