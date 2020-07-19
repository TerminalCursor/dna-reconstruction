package main

import (
	"fmt"
)

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

/*
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
