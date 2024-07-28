package selector

import (
	"github.com/SeaSkyThe/GeneticPhraseFinder/population"
)

type Selector struct{}

func NewSelector() *Selector {
	return &Selector{}
}

func (s *Selector) applySelection(pop population.Population, selectionFunc func(population.Population) population.Population) population.Population {
	return selectionFunc(pop)
}

func (s *Selector) ElitistSelection(pop population.Population) population.Population {
	return s.applySelection(pop, func(population population.Population) population.Population {
		// Elitism, choose the best half of the population
		population = population.OrderByFitness()
		return population[0 : len(population)/2]
	})
}
