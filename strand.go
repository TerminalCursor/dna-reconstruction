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
	strin := ""
	for _, b := range s.bases {
		strin += string(b.base)
	}
	return strin
}

func (s Strand) Complement() Strand {
	out := Strand{}
	for _, b := range s.bases {
		out.bases = append(out.bases, b.Complement())
	}
	return out
}

func (s Strand) Length() int {
	return len(s.bases)
}
