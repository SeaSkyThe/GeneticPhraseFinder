package individual

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
