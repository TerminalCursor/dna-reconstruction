package main

import (
	"fmt"
)

// Matched Scaffolding
// Datastructure that holds all of the staples as well as the binding index for the staple partition
type MatchedScaffold struct {
	scaffold Strand
	staples []Staple
	sbinds [][]int
}

// Get the overall length of a list of strands
func TotalLength(s []Strand) int {
	totalLength := 0
	for _, staple := range s {
		totalLength += staple.Length()
	}
	return totalLength
}

// Pretty Print Scaffolding with Bindings
func (matchedScaffold MatchedScaffold) MatchedString() string {
	outputString := fmt.Sprintf("%s\n", matchedScaffold.scaffold.Bases())
	for i:= 0; i < matchedScaffold.scaffold.Length(); i++ {
		didPrint := false
		for j := 0; j < len(matchedScaffold.staples); j++ {
			for k := 0; k < matchedScaffold.staples[j].Length(); k++ {
				if matchedScaffold.sbinds[j][k] <= i && i < (matchedScaffold.sbinds[j][k] + matchedScaffold.staples[j].pieces[k].Length()) && matchedScaffold.sbinds[j][k] != -1 {
					outputString += fmt.Sprintf("%c", matchedScaffold.staples[j].pieces[k].Bases()[i-matchedScaffold.sbinds[j][k]])
					didPrint = true
				}
			}
		}
		if !didPrint {
			outputString += fmt.Sprintf(" ")
		}
	}
	outputString += "\n"
	return outputString
}

// Score each bound staple strand
func (matched MatchedScaffold) Score() int {
	score := 0
	for i := 0; i < len(matched.staples); i++ {
		for idx, part := range matched.sbinds[i] {
			if part != -1 {
				score += matched.staples[i].pieces[idx].Length()
			}
		}
	}
	return score
}
