package reproductor

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/SeaSkyThe/GeneticPhraseFinder/individual"
	// "github.com/SeaSkyThe/GeneticPhraseFinder/mutator"
	"github.com/SeaSkyThe/GeneticPhraseFinder/population"
)

type Reproductor struct {
	totalFitness         float32
	multiPointCrossoverK int
}

func NewReproductor(totalFitness float32, multiPointCrossoverK int) *Reproductor {
	if multiPointCrossoverK < 1 {
		multiPointCrossoverK = 1
	}
	return &Reproductor{totalFitness: totalFitness, multiPointCrossoverK: multiPointCrossoverK}
}

func (r *Reproductor) Reproduce(pop population.Population, crossoverMethod func(parent1 *individual.Individual, parent2 *individual.Individual) (*individual.Individual, error)) (population.Population, error) {
	newPop := make(population.Population, len(pop))
	copy(newPop, pop)

	for i := 0; i < len(pop); i++ {
		parent1 := r.rouletteWheelSelection(pop)
		parent2 := r.rouletteWheelSelection(pop)
		child, err := crossoverMethod(parent1, parent2)
		if err != nil {
			return nil, fmt.Errorf("Error generating crossover: %s", err)
		}
		newPop = append(newPop, child)
	}

	return newPop, nil
}

func (r *Reproductor) SinglePointCrossover(pop population.Population, mutationMethod func(genes string) string) (population.Population, error) {
	return r.Reproduce(pop, func(parent1 *individual.Individual, parent2 *individual.Individual) (*individual.Individual, error) {
		if len(parent1.Genes) != len(parent2.Genes) {
			return nil, errors.New("Parents must have the same length")
		}

		var crossoverPoint int = len(parent1.Genes) / 2
		childGenes := parent1.Genes[0:crossoverPoint] + parent2.Genes[crossoverPoint:]
		childGenes = mutationMethod(childGenes)
		child := individual.NewIndividual(childGenes, 0)

		return child, nil
	})
}

func (r *Reproductor) MultiPointCrossover(pop population.Population, mutationMethod func(genes string) string) (population.Population, error) {
	return r.Reproduce(pop, func(parent1 *individual.Individual, parent2 *individual.Individual) (*individual.Individual, error) {
		if len(parent1.Genes) != len(parent2.Genes) {
			return nil, errors.New("Parents must have the same length")
		}

		if r.multiPointCrossoverK < 1 {
			return nil, errors.New("k must be greater than 0")
		}

		stepSize := len(parent1.Genes) / (r.multiPointCrossoverK + 1)

		childGenes := ""
		p1 := true
		for i := 0; i <= r.multiPointCrossoverK; i++ {
			start := i * stepSize
			end := start + stepSize

			if end > len(parent1.Genes) {
				end = len(parent1.Genes)
			}

			if p1 {
				childGenes += parent1.Genes[start:end]
			} else {
				childGenes += parent2.Genes[start:end]
			}
			p1 = !p1
		}

		// Handle the remaining part of the genes
		remainingStart := (r.multiPointCrossoverK + 1) * stepSize
		if remainingStart < len(parent1.Genes) {
			if p1 {
				childGenes += parent1.Genes[remainingStart:]
			} else {
				childGenes += parent2.Genes[remainingStart:]
			}
		}

		// Mutation
		childGenes = mutationMethod(childGenes)

		child := individual.NewIndividual(childGenes, 0)

		return child, nil
	})
}

// Other reproduction methods in the format:
// func (Population, func(string) string) (Population, error)

func (r *Reproductor) rouletteWheelSelection(pop population.Population) *individual.Individual {
	randValue := rand.Float32() * r.totalFitness
	accumulatedFitness := float32(0)

	for _, individual := range pop {
		accumulatedFitness += individual.Fitness
		if randValue <= accumulatedFitness {
			return individual
		}
	}

	// In case of errors, return the last individual
	return pop[len(pop)-1]
}
