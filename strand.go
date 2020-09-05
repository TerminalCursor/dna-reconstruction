package main

type Strand struct {
	bases []Nucleotide
}

func MakeStrand(bases string) Strand {
	strand := Strand{}
	for _, b := range bases {
		if b == 'A' || b == 'T' || b == 'G' || b == 'C' {
			strand.bases = append(strand.bases, MakeNucleotide(b))
		}
	}
	return strand
}

func (s Strand) Bases() string {
	bases := ""
	for _, b := range s.bases {
		bases += string(b.base)
	}
	return bases
}

func (s Strand) Complement() Strand {
	complement := Strand{}
	for _, b := range s.bases {
		complement.bases = append(complement.bases, b.Complement())
	}
	return complement
}

func (s Strand) Length() int {
	return len(s.bases)
}
