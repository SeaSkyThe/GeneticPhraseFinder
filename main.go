package main

import (
	"fmt"
	"time"

	"github.com/SeaSkyThe/GeneticPhraseFinder/mutator"
	"github.com/SeaSkyThe/GeneticPhraseFinder/population"
	"github.com/SeaSkyThe/GeneticPhraseFinder/reproductor"
	"github.com/SeaSkyThe/GeneticPhraseFinder/selector"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const POPULATION_SIZE = 150
const MUTATION_RATE = 0.015
const TARGET = "Olha que coisa mais linda, Mais cheia de graca, E ela menina, Que vem e que passa"

const MULTI_POINT_CROSSOVER_K = 4

const SCREEN_WIDTH = 1050
const SCREEN_HEIGHT = 700
const UPDATE_INTERVAL = time.Millisecond * 5
const MAX_HISTORY = 25

func main() {

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "This is my Genetic Algo Viz")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	population := population.GeneratePopulation(POPULATION_SIZE, len(TARGET))
	population.CalculatePopulationFitness(TARGET)

	var pause bool = true
	var framesCounter int = 0
	var lastUpdate time.Time
	var generation int = 0

	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.RayWhite)

	customFont := rl.LoadFontEx("./RobotoMonoNerdFont-Regular.ttf", 20, nil, 0)

	// customFont := rl.GetFontDefault()

	for rl.WindowShouldClose() == false {
		if rl.IsKeyPressed(rl.KeySpace) {
			pause = !pause
		}

		mostFit := population.GetMostFit()
		if !pause && mostFit.Genes != TARGET {
			now := time.Now()
			if now.Sub(lastUpdate) > UPDATE_INTERVAL {
				mostFit = population.GetMostFit()
				if mostFit == nil {
					break
				}

				// Selection
				selector := selector.NewSelector()

				// Reprodution + Mutation
				reproductor := *reproductor.NewReproductor(population.GetTotalFitness(), MULTI_POINT_CROSSOVER_K)
				mutator := *mutator.NewMutator(MUTATION_RATE)

				population.GenerateNextGeneration(selector.ElitistSelection,
					reproductor.OrderCrossover,
					mutator.RandomGenes,
					MUTATION_RATE)

				// Calculate Fitness
				population.CalculatePopulationFitness(TARGET)
				population.PrintPopulation()

				generation += 1
				lastUpdate = now
			}
		} else {
			framesCounter += 1
		}

		// Drawing data
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		mostFit = population.GetMostFit()
		if mostFit != nil {
			rl.DrawTextEx(customFont, fmt.Sprintf("Generation: %d", generation), rl.Vector2{X: 10, Y: 10}, 20, 2, rl.DarkGray)
			rl.DrawTextEx(customFont, fmt.Sprintf("Best Fitness: %.6f", mostFit.Fitness), rl.Vector2{X: 10, Y: 40}, 20, 2, rl.DarkGray)
			rl.DrawTextEx(customFont, fmt.Sprintf("Best %d Genes in Current Generation: ", MAX_HISTORY), rl.Vector2{X: 10, Y: 70}, 20, 2, rl.DarkGray)

			for i, ind := range population.OrderByFitness() {
				if i < MAX_HISTORY {
					yPosition := float32(95 + (i * 20))
					if i == 0 {
						if mostFit.Genes == TARGET {
							rl.DrawTextEx(customFont, fmt.Sprintf("[%2d]: %s", i+1, ind.Genes), rl.Vector2{X: 10, Y: yPosition}, 20, 2, rl.Green)
						} else {
							rl.DrawTextEx(customFont, fmt.Sprintf("[%2d]: %s", i+1, ind.Genes), rl.Vector2{X: 10, Y: yPosition}, 20, 2, rl.Black)
						}
					} else {
						rl.DrawTextEx(customFont, fmt.Sprintf("[%2d]: %s", i+1, ind.Genes), rl.Vector2{X: 10, Y: yPosition}, 20, 2, rl.DarkGray)
					}
				} else {
					break
				}
			}
		}

		// Pause handling

		if pause && ((framesCounter/30)%2) == 0 {
			rl.DrawText("PAUSED", int32(SCREEN_WIDTH/2-100), int32(SCREEN_HEIGHT/2-100), 30, rl.Gray)
		}

		rl.DrawFPS(int32(rl.GetScreenWidth())-100, 10)

		rl.DrawLine(0, SCREEN_HEIGHT-100, SCREEN_WIDTH, SCREEN_HEIGHT-100, rl.DarkGray)
		rl.DrawTextEx(customFont, fmt.Sprintf("Population size: %d", POPULATION_SIZE), rl.Vector2{X: 10, Y: SCREEN_HEIGHT - 90}, 20, 2, rl.DarkGray)
		rl.DrawTextEx(customFont, fmt.Sprintf("Mutation rate: %f", MUTATION_RATE), rl.Vector2{X: 10, Y: SCREEN_HEIGHT - 60}, 20, 2, rl.DarkGray)
		rl.DrawTextEx(customFont, fmt.Sprintf("Target: \"%s\"", TARGET), rl.Vector2{X: 10, Y: SCREEN_HEIGHT - 30}, 20, 2, rl.DarkGray)

		rl.EndDrawing()
	}
}
