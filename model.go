package main

type Nucleotide struct {
	base byte
}

type Strand struct {
	strand []Nucleotide
}

func MakeStrand(bases string) Strand {
	var outStrand Strand
	for i := 0; i < len(bases); i++ {
		if(bases[i] != 0x20) {
			outStrand.strand = append(outStrand.strand, Nucleotide{bases[i]})
		}
	}
	return outStrand
}

func (n Nucleotide) Complement() Nucleotide {
	var outN Nucleotide
	switch n.base {
		case 0x41:
			outN.base = 0x54
		case 0x54:
			outN.base = 0x41
		case 0x43:
			outN.base = 0x47
		case 0x47:
			outN.base = 0x43
		default:
			outN.base = 0x20
	}
	return outN
}

func (s Strand) Length() int {
	return len(s.strand)
}

func (s Strand) Complement() Strand {
	var outStrand Strand
	for i := 0; i < s.Length(); i++ {
		outStrand.strand = append(outStrand.strand, s.strand[i].Complement())
	}
	return outStrand
}

func (s Strand) Bases() string {
	base := ""
	for i := 0; i < s.Length(); i++ {
		base += string(s.strand[i].base)
	}
	return base
}

func (substrand Strand) Match(strand Strand) []int {
	var indices []int
	indices = append(indices, -1)
	for i := 0; i < strand.Length() - substrand.Length() + 1; i++ {
		for j := 0; j < substrand.Length(); j++ {
			if strand.strand[i + j] != substrand.strand[j].Complement() {
				break
			}
			if j + 1 == substrand.Length() {
				indices = append(indices, i)
			}
		}
	}
	return indices
}
