package main

import (
	"fmt"
)

/* https://en.wikipedia.org/wiki/ANSI_escape_code */

func main() {
	fmt.Printf("\033[91mDNA\033[0m \033[94mReconstruction\033[0m \033[33mv00.01\033[0m\n")
	fmt.Printf("%s\n%v\n", "ACTGCTGA", MakeStaple(MakeStrand("ACTGCTGA"), 4))
}
