package individual

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"

	"github.com/SeaSkyThe/GeneticPhraseFinder/mutator"
)

const EPSILON float32 = 0.001

type Individual struct {
	Genes   string
	Fitness float32
}

func (i *Individual) CalculateFitness(target string) {
	fitness := 0.0
	for j := range i.Genes {
		if i.Genes[j] == target[j] {
			fitness += 1.0
		}

	}
	(*i).Fitness = float32(fitness)/float32(len(target)) + EPSILON
}

func NewIndividual(genes string, fitness float32) *Individual {
	return &Individual{Genes: genes, Fitness: fitness}
}

type Population []*Individual

func (p Population) GetMostFit() *Individual {
	var best *Individual = nil
	for _, individual := range p {
		if best == nil || individual.Fitness > best.Fitness {
			best = individual
		}
	}
	return best
}

func (p *Population) GeneratePopulationFitness(target string) {
	for _, individual := range *p {
		(*individual).CalculateFitness(target)
	}
}

func (p *Population) GetPopulationOrderedByFitness() Population {
	// Sort population by fitness in descending order
	sort.SliceStable(*p, func(i, j int) bool {
		return (*p)[i].Fitness > (*p)[j].Fitness
	})

	return *p
}

func (p *Population) GenerateNextGeneration(mutationRate float32) error {
	var totalFitness float32 = 0.0
	for _, individual := range *p {
		totalFitness += individual.Fitness
	}

	// Sort population by fitness in descending order
	*p = (*p).GetPopulationOrderedByFitness()

	// Create a new generation
	nextGeneration := make(Population, 0, len(*p))

	// Elitism, choose the best half of the population
	for i := 0; i < len(*p)/2; i++ {
		nextGeneration = append(nextGeneration, (*p)[i])
	}

	// Generate the other half of the population using Crossover
	for i := len(*p) / 2; i < len(*p); i++ {
		parent1 := p.rouletteWheelSelection(totalFitness)
		parent2 := p.rouletteWheelSelection(totalFitness)
		child, err := Crossover(parent1, parent2, mutationRate)
		if err != nil {
			return fmt.Errorf("Error generating crossover: %s", err)
		}

        // Mutation
        mutator := *mutator.NewMutator(mutationRate)
        child.Genes = mutator.RandomGenes(child.Genes)

		nextGeneration = append(nextGeneration, child)
	}

	*p = nextGeneration

	return nil
}

func (p Population) rouletteWheelSelection(totalFitness float32) *Individual {
	randValue := rand.Float32() * totalFitness
	accumulatedFitness := float32(0)

	for _, individual := range p {
		accumulatedFitness += individual.Fitness
		if randValue <= accumulatedFitness {
			return individual
		}
	}

	// In case of errors, return the last individual
	return p[len(p)-1]
}

func (p Population) PrintPopulation() {
	for _, individual := range p {
		fmt.Println(individual.Genes, " | ", individual.Fitness)
	}
}

func GeneratePopulation(populationSize int, geneSize int) Population {
	var population Population = make([]*Individual, populationSize)
	for i := 0; i < populationSize; i++ {
		population[i] = NewIndividual(GenerateRandomGenes(geneSize), 0)
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

func Crossover(parent1 *Individual, parent2 *Individual, mutationRate float32) (*Individual, error) {
	if len(parent1.Genes) != len(parent2.Genes) {
		return nil, errors.New("Parents must have the same length")
	}

	var child *Individual = NewIndividual("", 0)
	var crossoverPoint int = len(parent1.Genes) / 2

    child.Genes = parent1.Genes[0:crossoverPoint] + parent2.Genes[crossoverPoint:]

    // Mutation
    mutator := *mutator.NewMutator(mutationRate)
    child.Genes = mutator.RandomGenes(child.Genes)

	return child, nil
}


