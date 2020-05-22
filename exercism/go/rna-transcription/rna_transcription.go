package strand

// ToRNA will transcribe given DNA strain to its corresponding RNA strain
func ToRNA(dna string) string {
	// Assuming the DNA sequence consist of nucleotide in bytes
	dnaSeq := []byte(dna)
	rnaSeq := make([]byte, len(dnaSeq))

	// The mapping of DNA nucleotide to RNA nucleotide
	rnaMapping := map[uint8]uint8{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	// Transcribing RNA from given DNA sequence
	for index, value := range dnaSeq {
		rnaSeq[index] = rnaMapping[value]
	}

	return string(rnaSeq)
}
