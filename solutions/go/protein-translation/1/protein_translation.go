package protein

import (
	"errors"
)

var (
	ErrStop        = errors.New("stop")
	ErrInvalidBase = errors.New("invalid codon")
)

var codonMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	var res []string
	for i := 0; i+3 <= len(rna); i += 3 {
		protein, err := FromCodon(rna[i : i+3])
		if err == ErrStop {
			return res, nil
		}
		if err != nil {
			return res, err
		}

		res = append(res, protein)
	}
	return res, nil
}

func FromCodon(codon string) (string, error) {
	protein, exists := codonMap[codon]
	if !exists {
		return "", ErrInvalidBase
	}
	if protein == "STOP" {
		return "", ErrStop
	}
	return protein, nil
}
