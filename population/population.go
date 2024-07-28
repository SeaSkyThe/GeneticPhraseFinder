package population

import (
	"errors"
	"fmt"
	"sort"

	"github.com/SeaSkyThe/GeneticPhraseFinder/individual"
	"github.com/SeaSkyThe/GeneticPhraseFinder/mutator"
)

type Population []*individual.Individual

func (p Population) GetMostFit() *individual.Individual {
	var best *individual.Individual = nil
	for _, individual := range p {
		if best == nil || individual.Fitness > best.Fitness {
			best = individual
		}
	}
	return best
}

func (p *Population) CalculatePopulationFitness(target string) {
	for _, individual := range *p {
		(*individual).CalculateFitness(target)
	}
}

func (p *Population) GetTotalFitness() float32 {
	var totalFitness float32 = 0.0
	for _, individual := range *p {
		totalFitness += individual.Fitness
	}
	return totalFitness
}

func (p *Population) GenerateNextGeneration(selectionMethod func(Population) Population,
	reproductionMethod func(Population, func(string) string) (Population, error),
	mutationMethod func(string) string,
	mutationRate float32) error {
	// Selection
	// Elitism, choose the best half of the population
	halfPop := selectionMethod(*p)

	fmt.Println(len(halfPop))

	// Crossover + Mutation
	nextGeneration, err := reproductionMethod(halfPop, mutationMethod)
	if err != nil {
		return err
	}

	*p = nextGeneration

	return nil
}

func (p Population) PrintPopulation() {
	for _, individual := range p {
		fmt.Println(individual.Genes, " | ", individual.Fitness)
	}
}

func (p Population) OrderByFitness() Population {
	// Sort population by fitness in descending order
	sort.SliceStable(p, func(i, j int) bool {
		return p[i].Fitness > p[j].Fitness
	})

	return p
}

// Auxiliary functions
func GeneratePopulation(populationSize int, geneSize int) Population {
	var population Population = make([]*individual.Individual, populationSize)
	for i := 0; i < populationSize; i++ {
		population[i] = individual.NewIndividual(GenerateRandomGenes(geneSize), 0)
	}
	return population
}

func GenerateRandomGenes(genesSize int) string {
	var genes string = ""
	for i := 0; i < genesSize; i++ {
		genes += string(mutator.RandomChar())
	}
	return genes
}

// Reproduction

func Crossover(parent1 *individual.Individual, parent2 *individual.Individual, mutationRate float32) (*individual.Individual, error) {
	if len(parent1.Genes) != len(parent2.Genes) {
		return nil, errors.New("Parents must have the same length")
	}

	var child *individual.Individual = individual.NewIndividual("", 0)
	var crossoverPoint int = len(parent1.Genes) / 2

	child.Genes = parent1.Genes[0:crossoverPoint] + parent2.Genes[crossoverPoint:]

	// Mutation
	mutator := *mutator.NewMutator(mutationRate)
	child.Genes = mutator.RandomGenes(child.Genes)

	return child, nil
}
