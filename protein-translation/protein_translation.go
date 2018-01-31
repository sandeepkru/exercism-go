package protein

var codonMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromCodon translates given codon to protien.
func FromCodon(input string) string {
	return codonMap[input]
}

// FromRNA translates given rna sequence to list of protiens.
func FromRNA(input string) []string {
	var result []string
	var protiens []string
	var substrs []string
	for i, j := 0, 3; j <= len(input); i, j = i+3, j+3 {
		substrs = append(substrs, input[i:j])
	}
	for _, cur := range substrs {
		if codonMap[cur] == "STOP" {
			break
		}
		protiens = append(protiens, codonMap[cur])
	}
	m := make(map[string]bool)
	for _, cur := range protiens {
		if exist, ok := m[cur]; exist || ok {
			continue
		}
		m[cur] = true
		result = append(result, cur)
	}
	return result
}
