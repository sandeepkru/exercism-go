package strand

import "log"

var m = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

// ToRNA converts given DNA to RNA sequence.
func ToRNA(dna string) string {
	var rna string
	for _, char := range dna {
		c, ok := m[string(char)]
		if !ok {
			log.Fatalf("DNA to RNA mapping not found for %v", char)
		}
		rna += c
	}

	return rna
}

// go test -run ^NOTHING -bench 'BenchmarkRNATranscription'
// goos: darwin
// goarch: amd64
// BenchmarkRNATranscription-8   	 2000000	       883 ns/op
// PASS
// ok  	_/Users/sandeep/practice/go-exercism/go/rna-transcription	2.689s
