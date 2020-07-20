package main

import (
	"fmt"
	"sync"
	"time"
)

func TestWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	fmt.Printf("\033[1;1H\033[2J\033[91mDNA\033[0m \033[94mReconstruction\033[0m \033[33mv00.03\033[0m\n")
	// Nucleotide Length for Staple Partitions
	NUCLEOTIDE_LENGTH := 6

	// Get Staple Strands from staples.txt
	for _, s := range ReadLines("sample/staples.txt") {
		fmt.Printf("%s\n", s)
	}

	fmt.Printf("%c\n", MakeNucleotide('A').base)
	fmt.Printf("%s\n", MakeStrand("ACTG CAGT").Bases()[:NUCLEOTIDE_LENGTH])
	strand := MakeStrand("ACTG CAGT")
	fmt.Printf("%s\n", strand.Bases())
	fmt.Printf("%s\n", strand.Complement().Bases())
	fmt.Printf("%s\n", strand.Bases())

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go TestWorker(i, &wg)
	}
	wg.Wait()

	scaffold := Scaffold{
		MakeStrand("AGTCGTCA"),
		[][]Strand{
			[]Strand{
				MakeStrand("TCAG"),
				MakeStrand("CAGT"),
			},
		},
		[][]int{
			[]int{
				0,
				4,
			},
		},
	}
	fmt.Printf("%s\n", scaffold.scaffold.Bases())
	for sidx, staple := range scaffold.staples {
		for idx, strand := range staple {
			fmt.Printf("%v %s\n", scaffold.bonds[sidx][idx], strand.Bases())
		}
	}
	fmt.Printf("%s", scaffold.MatchString())
	fmt.Printf("%v\n", scaffold.MatchStrand(MakeStrand("GT")))
}
