# GeneticPhraseFinder
A genetic algorithm to match a desired phrase implemented in Go.

The visualization is done using [raylib-go](https://github.com/gen2brain/raylib-go).

## How to run

```bash
go run main.go
```

## How it works

The genetic algorithm is based on the following principles:

- The population is made of individuals, each with a genetic code.
- The fitness of each individual is calculated based on the number of matching characters between the genetic code and the target phrase.
- The fittest individuals are selected to create the next generation (half of the current population.
- The next generation is created by performing crossover between the fittest individuals.
  - The crossover point is always half of the genetic code length.
  - The genetic code of each individual is mutated with a certain probability, to introduce randomness, and to avoid getting stuck in local optimal.
    - The mutation can happen at any position in the genetic code.
- The process is repeated until the target phrase is found.

## References

- [Genetic Algorithms](https://en.wikipedia.org/wiki/Genetic_algorithm)
- Genetic Algorithm - A Literature Review - DOI : *10.1109/COMITCon.2019.8862255*
