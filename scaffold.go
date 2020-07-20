package main

import (
	"sync"
)

type Scaffold struct {
	scaffold Strand
	staples [][]Strand
	bonds [][]int
}

// Get a pretty-print output of Scaffold and the bonded staples
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

// Find all of the binding sites on the scaffold for a strand
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
				// Wait for all partial staple strand thread checks to come back
				wg.Wait()
			}
			// If nothing is bonded, add that as an available bond site
			if available {
				bondSites = append(bondSites, i)
			}
		}
	}
	return bondSites
}

// Check if bonding site is in list of valid bonding sites
func InIntList(val int, list []int) bool {
	isIn := false
	for _, i := range list {
		if val == i {
			isIn = true
		}
	}
	return isIn
}

// Bond a staple to the scaffold, only if it will not collide with existing bonds
func (sc Scaffold) BondStaple(strands []Strand, bonds []int) Scaffold {
	isValid := true
	for sidx, strand := range strands {
		validBonds := sc.MatchStrand(strand)
		isValid = isValid && InIntList(bonds[sidx], validBonds)
	}
	if isValid {
		sc.staples = append(sc.staples, strands)
		sc.bonds = append(sc.bonds, bonds)
	}
	return sc
}
