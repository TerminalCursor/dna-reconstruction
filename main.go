package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
)

func PrintMatch(scaffold, staplePart Strand, offset int) {
	fmt.Printf("%s\n", scaffold.Bases())
	for i := 0; i < offset + staplePart.Length(); i++ {
		if i >= offset { fmt.Printf("%c", staplePart.Bases()[i-offset])
		} else {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
}

func GetPermutations(input [][]int, prep []int, interference []int, WINDOW_SIZE int) [][]int {
	var out [][]int
	if(len(input) > 1) {
		if len(input[0]) > 0 {
			for i := 0; i < len(input[0]); i++ {
				validPath := true
				for j := 0; j < len(prep); j++ {
					if prep[j] <= input[0][i] && input[0][i] < (prep[j] + WINDOW_SIZE) && prep[j] != -1 { validPath = false
						break
					}
				}
				for j := 0; j < len(interference); j++ {
					if interference[j] <= input[0][i] && input[0][i] < (interference[j] + WINDOW_SIZE) && interference[j] != -1 {
						validPath = false
						break
					}
				}
				if validPath {
					perms := GetPermutations(input[1:], append(prep, input[0][i]), interference, WINDOW_SIZE)
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
			for j := 0; j < len(interference); j++ {
				if interference[j] <= input[0][i] && input[0][i] < (interference[j] + WINDOW_SIZE) && interference[j] != -1 {
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

func GetStaples(input [][]Staple, prep []Staple) [][]Staple {
	var out [][]Staple
	if(len(input) > 1) {
		if len(input[0]) > 0 {
			for i := 0; i < len(input[0]); i++ {
				perms := GetStaples(input[1:], append(prep, input[0][i]))
				for j := 0; j < len(perms); j++ {
					out = append(out, perms[j])
				}
			}
		} else {
			fmt.Printf("EMPTY MATCH DETECTED\nTRIVIAL MATCH '-1' NOT INCLUDED\n")
			os.Exit(1)
		}
	} else {
		for i := 0; i < len(input[0]); i++ {
			out = append(out, append(prep, input[0][i]))
		}
	}
	return out
}

func BindingJoin(bindingSites [][]int) []int {
	var output []int
	for i := 0; i < len(bindingSites); i++ {
		for _, val := range bindingSites[i] {
			output = append(output, val)
		}
	}
	return output
}

/* https://en.wikipedia.org/wiki/ANSI_escape_code */

func main() {
	fmt.Printf("\033[91mDNA\033[0m \033[94mReconstruction\033[0m \033[33mv00.01\033[0m\n")
	/* MULTI STAPLES */
	WINDOW_SIZE := 8
	var staple_strands []Strand
	/*
	FILE READING
	*/
	file, err := os.Open("staples.txt")
	if err != nil {
		log.Fatalf("Failed opening file: %s\n", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		staple_strands = append(staple_strands, MakeStrand(scanner.Text()))
	}

	file.Close()
	/*
	END FILE READING
	*/
	// Generate all possible staple subsections of WINDOWSIZE
	var staple_options [][]Staple
	for idx, staple_strand := range staple_strands {
		staple_options = append(staple_options, []Staple{})
		if int(staple_strand.Length() / 2) >= WINDOW_SIZE {
			for i := 0; i < WINDOW_SIZE; i++ {
				staple_options[idx] = append(staple_options[idx], MakeStaple(MakeStrand(staple_strand.Bases()[i:]), WINDOW_SIZE))
			}
			fmt.Printf("Staple Chunk: %d of %d\n", idx + 1, len(staple_strands))
		} else {
			fmt.Printf("Cannot generate 2 staples from given sequence at %d windowsize\n", WINDOW_SIZE)
			os.Exit(1)
		}
	}
	cut_length := TotalLength(staple_strands)
	if cut_length > m13mp18f.Length() {
		cut_length = m13mp18f.Length()
	}
	m13mp18 := MakeStrand(m13mp18f.Bases()[:cut_length]).Reverse()
	var topMatch []MatchedScaffold
	var sorted_staples = GetStaples(staple_options, []Staple{})
	for idx, staples := range sorted_staples {
		fmt.Printf("Round: %d of %d\n", idx + 1, len(sorted_staples))
		var matchedScaffolds []MatchedScaffold = []MatchedScaffold{MatchedScaffold{m13mp18, []Staple{}, [][]int{}}}
		var nextMatchedScaffolds []MatchedScaffold
		for _, staple := range staples {
			nextMatchedScaffolds = []MatchedScaffold{}
			var staplePartitionMatches [][]int
			for i := 0; i < staple.Length(); i++ {
				staplePartitionMatches = append(staplePartitionMatches, staple.pieces[i].Match(m13mp18))
				//fmt.Printf("%s %v\n", staple.pieces[i].Bases(), staplePartitionMatches[len(staplePartitionMatches)-1])
			}
			//fmt.Printf("%v\n", staplePartitionMatches)
			var possibleCombinations [][]int
			for _, ms := range matchedScaffolds {
				if len(ms.staples) == 0 {
					possibleCombinations = GetPermutations(staplePartitionMatches, []int{}, []int{}, WINDOW_SIZE)
				} else {
					possibleCombinations = GetPermutations(staplePartitionMatches, []int{}, BindingJoin(ms.sbinds), WINDOW_SIZE)
				}
				for i := 0; i < len(possibleCombinations); i++ {
					nextMatchedScaffolds = append(nextMatchedScaffolds, MatchedScaffold{m13mp18, append(ms.staples, staple), append(ms.sbinds, possibleCombinations[i])})
				}
			}
			matchedScaffolds = nextMatchedScaffolds
		}
		topScore := 0
		for _, matched := range matchedScaffolds {
			if topScore < matched.Score() {
				topScore = matched.Score()
			}
		}
		for _, matched := range matchedScaffolds {
			if topScore == matched.Score() {
				topMatch = append(topMatch, matched)
			}
		}
	}
	for _, matched := range topMatch {
		fmt.Printf("Match Score: %d\n%s", matched.Score(), matched.MatchedString())
	}
}
