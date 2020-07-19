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
	// Make Staple Piece that is smaller than partition size
	/*
	if s.Length() % partition != 0 {
		staple.pieces = append(staple.pieces, MakeStrand(s.Bases()[FullStrandsLength * partition:]))
	}
	*/
	return staple
}

func (s Staple) Length() int {
	return len(s.pieces)
}

func (s Staple) TotalLength() int {
	totalLength := 0
	for i := 0; i < s.Length(); i++ {
		totalLength += s.pieces[i].Length()
	}
	return totalLength
}

func (s Staple) Bases() string {
	outString := ""
	for _,strand := range s.pieces {
		outString += strand.Bases()
	}
	return outString
}
