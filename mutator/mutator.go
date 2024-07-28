package mutator

import (
	"math/rand"
)

const CHARSET string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 ,"

type Mutator struct {
	mutationRate float32
}

func NewMutator(mutationRate float32) *Mutator {
	return &Mutator{mutationRate: mutationRate}
}

func (m *Mutator) applyMutation(genes string, mutationFunc func(*[]rune, int)) string {
	genesRunes := []rune(genes)
	for i := 0; i < len(genesRunes); i++ {
		if rand.Float32() <= m.mutationRate {
			mutationFunc(&genesRunes, i)
		}
	}
	return string(genesRunes)
}

func (m *Mutator) RandomGenes(genes string) string {
	return m.applyMutation(genes, func(genesRunes *[]rune, i int) {
		(*genesRunes)[i] = RandomChar()
	})
}

func (m *Mutator) SwapGenesPositions(genes string) string {
	return m.applyMutation(genes, func(genesRunes *[]rune, i int) {
		*genesRunes = SwapGenes(*genesRunes, i)
	})
}

var charsetRunes = []rune(CHARSET)

func RandomChar() rune {
	return charsetRunes[rand.Intn(len(charsetRunes))]
}

func SwapGenes(genes []rune, i int) []rune {
	size := len(genes)
	if size <= 1 {
		return genes
	}

	randPos := rand.Intn(size)
	for randPos == i {
		randPos = rand.Intn(size)
	}
	(genes)[i], (genes)[randPos] = (genes)[randPos], (genes)[i]

	return genes
}
