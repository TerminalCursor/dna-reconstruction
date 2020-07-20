package main

type Nucleotide struct {
	base rune
	pos [3]float64
}

func MakeNucleotide(base rune) Nucleotide {
	nucleotide := Nucleotide{base, [3]float64{0,0,0}}
	return nucleotide
}

func (n Nucleotide) Complement() Nucleotide {
	out := n
	switch out.base {
	case 0x41:
		out.base = 0x54
	case 0x54:
		out.base = 0x41
	case 0x43:
		out.base = 0x47
	case 0x47:
		out.base = 0x43
	}
	return out
}
