package main

import (
	"sync"
)

type Scaffold struct {
	scaffold Strand
	staples [][]Strand
	bonds [][]int
}

func (s Scaffold) MatchString() string {
	out := s.scaffold.Bases() + "\n"
	for i := 0; i < s.scaffold.Length(); i++ {
		test := false
		for sidx, staple := range s.staples {
			for idx, strand := range staple {
				if s.bonds[sidx][idx] <= i && i < s.bonds[sidx][idx] + strand.Length() {
					out += string(strand.Bases()[i - s.bonds[sidx][idx]])
					test = true
				}
			}
		}
		if !test {
			out += " "
		}
	}
	out += "\n"
	return out
}

func (sc Scaffold) MatchStrand(s Strand) []int {
	var bondSites []int
	for i := 0; i < sc.scaffold.Length() - s.Length() + 1; i++ {
		// Check to see if the strand can bond at every available position on the strand
		if sc.scaffold.Bases()[i:i+s.Length()] == s.Complement().Bases() {
			// If it can, check to see if anything is bonded in that range already
			available := true
			var wg sync.WaitGroup
			for sidx, sbond := range sc.bonds {
				for idx, bint := range sbond {
					wg.Add(1)
					// Check each staple on its own thread
					go func() {
						defer wg.Done()
						for j := bint; j < bint + sc.staples[sidx][idx].Length(); j ++ {
							if i <= j && j < i + s.Length() {
								available = false
							}
						}
					}()
				}
			}
			// Wait for all staple thread checks to come back
			wg.Wait()
			// If nothing is bonded, add that as an available bond site
			if available {
				bondSites = append(bondSites, i)
			}
		}
	}
	return bondSites
}
