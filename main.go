package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"sort"
)

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
	fmt.Printf("\033[91mDNA\033[0m \033[94mReconstruction\033[0m \033[33mv00.02\033[0m\n")
	/* MULTI STAPLES */
	WINDOW_SIZE := 6
	var staple_strands []Strand
	/*
	FILE READING - GET STAPLE STRANDS
	*/
	file, err := os.Open("staples.txt")
	if err != nil {
		log.Fatalf("Failed opening file: %s\n", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	// Each line is a staple strand
	for scanner.Scan() {
		staple_strands = append(staple_strands, MakeStrand(scanner.Text()))
	}

	file.Close()
	/*
	END FILE READING
	*/

	// Generate Scaffold
	cut_length := TotalLength(staple_strands)
	if cut_length > m13mp18f.Length() {
		cut_length = m13mp18f.Length()
	}
	m13mp18 := MakeStrand(m13mp18f.Bases()[:cut_length]).Reverse()

	// Match each strand and keep highest scores
	//var topScoredMatches []MatchedScaffold
	var matchedScaffolds []MatchedScaffold
	matchedScaffolds = []MatchedScaffold{MatchedScaffold{m13mp18, []Staple{}, [][]int{}}}
	var nextMatchedScaffolds []MatchedScaffold
	for sidx, staple_strand := range staple_strands {
		nextMatchedScaffolds = []MatchedScaffold{}
		fmt.Printf("Round: %d of %d\n", sidx + 1, len(staple_strands))
		// Generate possible chunks for a given staple
		var stapleOptions []Staple
		if int(staple_strand.Length() / 2) >= WINDOW_SIZE {
			for i := 0; i < WINDOW_SIZE; i++ {
				var generated_staple_strands []Strand
				// Get any piece of strand before the WINDOW_SIZE'd chunks
				if i != 0 {
					generated_staple_strands = append(generated_staple_strands, MakeStrand(staple_strand.Bases()[:i]))
				}
				// Get all of the WINDOW_SIZE chunks with 'i' offset
				j := i
				for ; j + WINDOW_SIZE <= staple_strand.Length(); j += WINDOW_SIZE {
					generated_staple_strands = append(generated_staple_strands, MakeStrand(staple_strand.Bases()[j:j+WINDOW_SIZE]))
				}
				// Get any remnants
				if (j + 1) <= staple_strand.Length() {
					generated_staple_strands = append(generated_staple_strands, MakeStrand(staple_strand.Bases()[j:]))
				}
				stapleOptions = append(stapleOptions, Staple{generated_staple_strands})
			}
		} else {
			fmt.Printf("Cannot generate 2 staples from given sequence at %d windowsize\n", WINDOW_SIZE)
			os.Exit(1)
		}
		// Finished Generating All Possible Staple Chunks for this Staple
		topScores := []int{0, 0, 0}
		// Match All Possible Staple Chunks to Scaffold
		for idx, staple := range stapleOptions {
			fmt.Printf("Staple: %d of %d\n", idx + 1, len(stapleOptions))
			var staplePartitionMatches [][]int
			for i := 0; i < staple.Length(); i++ {
				if staple.pieces[i].Length() != WINDOW_SIZE {
					staplePartitionMatches = append(staplePartitionMatches, []int{-1})
				} else {
					staplePartitionMatches = append(staplePartitionMatches, staple.pieces[i].Match(m13mp18))
				}
			}
			//for _, piece := range staple.pieces {
			//	fmt.Printf("%s ", piece.Bases())
			//}
			//fmt.Printf("%v\n", staplePartitionMatches)
			var possibleCombinations [][]int
			for ijk, matched := range matchedScaffolds {
				//nextMatchedScaffolds = append(nextMatchedScaffolds, matched)
				if len(matched.staples) == 0 {
					possibleCombinations = GetPermutations(staplePartitionMatches, []int{}, []int{}, WINDOW_SIZE)
				} else {
					possibleCombinations = GetPermutations(staplePartitionMatches, []int{}, BindingJoin(matched.sbinds), WINDOW_SIZE)
				}
				for i := 0; i < len(possibleCombinations); i++ {
					matchedScaffold := MatchedScaffold{m13mp18, append(matched.staples, staple), append(matched.sbinds, possibleCombinations[i])}
					if topScores[2] <= matchedScaffold.Score() {
						nextMatchedScaffolds = append(nextMatchedScaffolds, matchedScaffold)
					}
					if matchedScaffold.Score() > topScores[0] {
						topScores[2] = topScores[1]
						topScores[1] = topScores[0]
						topScores[0] = matchedScaffold.Score()
					} else if matchedScaffold.Score() > topScores[1] {
						topScores[2] = topScores[1]
						topScores[1] = matchedScaffold.Score()
					} else if matchedScaffold.Score() > topScores[2] {
						topScores[2] = matchedScaffold.Score()
					}
					if ijk % 100 == 0 {
					//if ijk % 100 == 0 || len(nextMatchedScaffolds) > 3500 {
						// Sort by score (highest to lowest)
						sort.Slice(nextMatchedScaffolds, func(i, j int) bool {
							return nextMatchedScaffolds[i].Score() >= nextMatchedScaffolds[j].Score()
						})
						MAX := 1500
						if len(nextMatchedScaffolds) < MAX {
							MAX = len(nextMatchedScaffolds)
						}
						// Clean up matches to MAX+100 max entries
						nextMatchedScaffolds = nextMatchedScaffolds[:MAX]
					}
				}
			}
		}
		fmt.Printf("=== ROUND %d RESULTS  ===\n", sidx + 1)
		if len(nextMatchedScaffolds) != 0 {
			//matchedScaffolds = []MatchedScaffold{}
			for _, ms := range nextMatchedScaffolds {
				if ms.Score() == topScores[0] || ms.Score() == topScores[1] || ms.Score() == topScores[2] {
					matchedScaffolds = append(matchedScaffolds, ms)
				}
			//	if ms.Score() == topScores[0] {
			//		fmt.Printf("%s\n", ms.MatchedString())
			//	}
			}
			sort.Slice(matchedScaffolds, func(i, j int) bool {
				return matchedScaffolds[i].Score() >= matchedScaffolds[j].Score()
			})
			MAX := 1000
			if len(matchedScaffolds) < MAX {
				MAX = len(matchedScaffolds)
			}
			matchedScaffolds = matchedScaffolds[:MAX]
			fmt.Printf("Total Scaffolds Found: %d\n", len(matchedScaffolds))
		} else {
			fmt.Printf("Total Scaffolds: %d\n", len(matchedScaffolds))
			fmt.Printf("No New Scaffolds Found for this Staple\n")
		}
		fmt.Printf("=== ROUND %d FINISHED ===\n", sidx + 1)
	}
	for idx, ms := range matchedScaffolds {
		if idx < 10 {
			fmt.Printf("%s", ms.MatchedString())
			fmt.Printf("%d\n", ms.Score())
		}
	}
}
