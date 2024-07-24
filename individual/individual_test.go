package individual

import (
	"reflect"
	"testing"
)

func TestGenerateRandomGenes(t *testing.T) {
	genes := GenerateRandomGenes(10)
	if len(genes) != 10 {
		t.Errorf("Expected length of 10, got %d", len(genes))
	}
    if reflect.TypeOf(genes) != reflect.TypeOf("") {
        t.Errorf("Expected type of string, got %s", reflect.TypeOf(genes))
    }
}

func TestGeneratePopulation(t *testing.T) {
	population := GeneratePopulation(10, 10)
	if len(population) != 10 {
		t.Errorf("Expected length of 10, got %d", len(population))
	}
}

func TestCrossover(t *testing.T) {
	parent1 := NewIndividual("abc123", 0)
	parent2 := NewIndividual("456def", 0)
	child, err := Crossover(parent1, parent2, 0.0)

    if err != nil {
        t.Errorf("Expected no error, got %s", err)
    }

	if child.Genes != "abcdef" {
		t.Errorf("Expected child.Genes to be \"abcdef\", got %s", child.Genes)
	}
}

func TestCrossoverDifferentLength(t *testing.T) {
	parent1 := NewIndividual("abc123", 0)
	parent2 := NewIndividual("456", 0)
	_, err := Crossover(parent1, parent2, 0.0)

    if err == nil {
        t.Errorf("Expected error, got nil")
    }
}
