package reproductor

import (
	"testing"

	"github.com/SeaSkyThe/GeneticPhraseFinder/individual"
	"github.com/SeaSkyThe/GeneticPhraseFinder/mutator"
	"github.com/SeaSkyThe/GeneticPhraseFinder/population"
)

func TestSinglePointCrossover(t *testing.T) {
	ind1 := individual.NewIndividual("abcde", 0)
	ind2 := individual.NewIndividual("fghij", 0)
	pop := population.Population{ind1, ind2}

	mutator := *mutator.NewMutator(0.0)
	newPop, err := NewReproductor(pop.GetTotalFitness()).SinglePointCrossover(pop, mutator.RandomGenes)

	if err != nil {
		t.Errorf("Error generating crossover: %s", err)
	}

	if len(pop) != 2 {
		t.Errorf("Population size must be 2, got %d", len(pop))
	}

	if len(newPop) != 4 {
		t.Errorf("Population size must be 4, got %d", len(newPop))
	}

}
