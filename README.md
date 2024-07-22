# GeneticPhraseFinder

A genetic algorithm to match a desired phrase implemented in Go.

The visualization is done using [raylib-go](https://github.com/gen2brain/raylib-go).

I made this algorithm just to have fun and took me like 4 hours to implement it, if you find some nasty code, please let me know.

## Example


https://github.com/user-attachments/assets/4352ab25-a6ab-491a-b7a9-1368e326fc1e



## How to run

```bash
go run main.go
```

After running press SPACE to unpause/pause the execution.

## Configuration

The following constants can be changed to change the behavior of the genetic algorithm:

```go
const POPULATION_SIZE = 800
const MUTATION_RATE = 0.03

const SCREEN_WIDTH = 1050
const SCREEN_HEIGHT = 700
const UPDATE_INTERVAL = time.Millisecond * 5 // Screen update interval
const MAX_HISTORY = 25
const TARGET = "Olha que coisa mais linda, Mais cheia de graca, E ela menina, Que vem e que passa"
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
- Genetic Algorithm - A Literature Review - DOI : _10.1109/COMITCon.2019.8862255_
