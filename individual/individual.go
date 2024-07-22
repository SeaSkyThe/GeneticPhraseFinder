package individual

import (
	"fmt"
	"math/rand"
	"sort"
)

const EPSILON float32 = 0.001
const CHARSET string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 ,"

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
	if best == nil {
		return nil
	}
	return best
}

func (p *Population) GeneratePopulationFitness(target string) {
	for _, individual := range *p {
		// fmt.Println("Generating fitness for: ", individual.Genes)
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
func (p *Population) GenerateNextGeneration(mutationRate float32) {

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
		child := Crossover(parent1, parent2, mutationRate)
		Mutate(child, mutationRate)
		nextGeneration = append(nextGeneration, child)
	}

	*p = nextGeneration

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

func RandomChar() rune {
	charset := []rune(CHARSET)
	return charset[rand.Intn(len(charset))]
}

func GenerateRandomGenes(genesSize int) string {
	var genes string = ""
	for i := 0; i < genesSize; i++ {
		genes += string(RandomChar())
	}
	return genes
}

func Crossover(parent1 *Individual, parent2 *Individual, mutationRate float32) *Individual {
	var child *Individual = NewIndividual("", 0)
	var crossoverPoint int = len(parent1.Genes) / 2

	child.Genes = parent1.Genes[0:crossoverPoint] + parent2.Genes[crossoverPoint:]
	Mutate(child, mutationRate)

	return child
}

func Mutate(individual *Individual, mutationRate float32) {
	genes := []rune(individual.Genes)
	for i := 0; i < len(genes); i++ {
		if rand.Float32() <= mutationRate {
			genes[i] = RandomChar()
			(*individual).Genes = string(genes)

		}
	}
}
