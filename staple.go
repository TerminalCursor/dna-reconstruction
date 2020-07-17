package main

type Staple struct {
	pieces []Strand
}

func MakeStaple(s Strand, partition int) Staple {
	FullStrandsLength := int(s.Length() / partition)
	var staple Staple
	for i := 0; i < FullStrandsLength; i++ {
		staple.pieces = append(staple.pieces, MakeStrand(s.Bases()[i*partition:(i+1)*partition]))
	}
	if s.Length() % partition != 0 {
		staple.pieces = append(staple.pieces, MakeStrand(s.Bases()[FullStrandsLength * partition:]))
	}
	return staple
}
