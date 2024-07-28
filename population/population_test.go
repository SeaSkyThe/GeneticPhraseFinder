package population

import "testing"

func TestGeneratePopulation(t *testing.T) {
	population := GeneratePopulation(10, 10)
	if len(population) != 10 {
		t.Errorf("Expected length of 10, got %d", len(population))
	}

    if len(population[0].Genes) != 10 {
        t.Errorf("Expected length of 10, got %d", len(population[0].Genes))
    }
}

