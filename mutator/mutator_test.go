package mutator

import (
	"testing"
)

func TestSwapGenesPositions(t *testing.T) {
	genes := "abcdef"

	geneRunes := []rune(genes)
	SwapGenes(&geneRunes, 1)
	result := string(geneRunes)

	if result == "abcdef" {
		t.Errorf("Expected result to be != \"abcdef\", got %s", result)
	}
}

func TestRandomChar(t *testing.T) {
	char := RandomChar()

	found := false
	for _, c := range CHARSET {
		if c == char {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected char to be in charset, got %c", char)
	}
}
