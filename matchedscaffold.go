package main

import (
	"fmt"
)

// Future version to include multiple staples
type MatchedScaffold struct {
	scaffold Strand
	staples []Staple
	sbinds [][]int
}

func TotalLength(s []Staple) int {
	totalLength := 0
	for _,staple := range s {
		totalLength += staple.TotalLength()
	}
	return totalLength
}

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

func (matched MatchedScaffold) Score() int {
	score := TotalLength(matched.staples)
	for i := 0; i < len(matched.staples); i++ {
		for idx, part := range matched.sbinds[i] {
			if part == -1 {
				score -= matched.staples[i].pieces[idx].Length()
			}
		}
	}
	return score
}

/* OLD 2.0 VERSION */
/*
type MatchedScaffold struct {
	scaffold Strand
	staple Staple
	binds []int
}

func (matchedScaffold MatchedScaffold) MatchedString() string {
	outputString := fmt.Sprintf("%s\n", matchedScaffold.scaffold.Bases())
	for i:= 0; i < matchedScaffold.scaffold.Length(); i++ {
		didPrint := false
		for j := 0; j < matchedScaffold.staple.Length(); j++ {
			if matchedScaffold.binds[j] <= i && i < (matchedScaffold.binds[j] + matchedScaffold.staple.pieces[j].Length())  && !didPrint && matchedScaffold.binds[j] != -1{
				outputString += fmt.Sprintf("%c", matchedScaffold.staple.pieces[j].Bases()[i-matchedScaffold.binds[j]])
				didPrint = true
			}
		}
		if !didPrint {
			outputString += fmt.Sprintf(" ")
		}
	}
	outputString += "\n"
	return outputString
}
*/

/* OLD VERSION
type MatchedScaffold struct {
	scaffold Strand
	staple Staple
	binds []int
}

func MakeMatchedScaffold(scaff Strand, staple Staple, bindingPoints []int) {
	var ms MatchedScaffold
	ms.scaffold = scaff
	ms.staple = staple
	ms.binds = bindingPoints
}

func (matchedScaffold MatchedScaffold) PrintMatched() {
	fmt.Printf("%s\n", matchedScaffold.scaffold.Bases())
	for i:= 0; i < matchedScaffold.scaffold.Length(); i++ {
		didPrint := false
		for j := 0; j < matchedScaffold.staple.Length(); j++ {
			if matchedScaffold.binds[j] <= i && i < (matchedScaffold.binds[j] + matchedScaffold.staple.pieces[j].Length())  && !didPrint && matchedScaffold.binds[j] != -1{
				fmt.Printf("%c", matchedScaffold.staple.pieces[j].Bases()[i-matchedScaffold.binds[j]])
				didPrint = true
			}
		}
		if !didPrint {
			fmt.Printf(" ")
		}
	}
	fmt.Println()
}
*/
