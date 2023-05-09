package main

import (
	"fmt"
	"strings"
)

var CurrentGeneration int = 0

const (
	//chaine a devine
	EXPECTED_STR = "LA PHRASE A deviner ET C'est PAS SIMPLE JE PENSE"
	//pop max, generation maximum
	POPULATION_COUNT     = 100
	GENERATION_COUNT_MAX = 100000

	/* --- pourcentages ---*/
	//chance de muter
	CHANCE_TO_MUTATE = 0.4
	//pourcentage de population grade garder
	GRADED_RETAIN_PERCENT = 0.2
	//chance de retenir des non grade
	CHANCE_RETAIN_NONGRATED = 0.05

	GRADED_INDIVIDUAL_RETAIN_COUNT = int(POPULATION_COUNT * GRADED_RETAIN_PERCENT)

	LENGTH_OF_EXPECTED_STR = len(EXPECTED_STR)

	MIDDLE_LENGTH_OF_EXPECTED_STR = LENGTH_OF_EXPECTED_STR / 2

	ASCII_LETTERS        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ALLOWED_SPECIAL_CHAR = " !'."

	ALLOWED_CHARMAP = ASCII_LETTERS + ALLOWED_SPECIAL_CHAR

	MAXIMUM_FITNESS = LENGTH_OF_EXPECTED_STR
)

func main() {
	// Create a population and compute starting grade
	population := getRandomPop()
	averageGrade := getAverageFitness(population)
	fmt.Printf("Starting grade: %.2f / %d\n", averageGrade, MAXIMUM_FITNESS)

	// Make the population evolve
	i := 0
	var solution []Individual
	for solution == nil && i < GENERATION_COUNT_MAX {
		// Evolve population
		population, averageGrade, solution = evolvePopulation(population)

		fmt.Println(i)
		// Print progress every 256 generations
		if i&255 == 255 {
			fmt.Printf("Current grade: %.2f / %d (%d generation)\n", averageGrade, MAXIMUM_FITNESS, i)
		}

		// Increment generation counter
		i++
		fmt.Printf("Current grade: %.2f / %d (%d generation)\n", averageGrade, MAXIMUM_FITNESS, i)

	}

	// Create and render fitness evolution chart
	/* 	lineChart := pygal.Line{
	   		ShowDots:   false,
	   		ShowLegend: false,
	   		Title:      "Fitness evolution",
	   		XTitle:     "Generations",
	   		YTitle:     "Fitness",
	   	}
	   	for _, avg := range logAvg {
	   		lineChart.Add("Fitness", avg)
	   	}
	   	lineChart.RenderToSVG("fitness_evolution.svg") */

	// Print the final stats
	averageGrade = getAverageFitness(population)
	fmt.Printf("Final grade: %.2f / %d\n", averageGrade, MAXIMUM_FITNESS)

	// Print the solution
	if solution != nil {
		fmt.Printf("Solution found (%d times) after %d generations.\n", len(solution), i)
	} else {
		fmt.Printf("No solution found after %d generations.\n", i)
		fmt.Println("- Last population was:")
		for i, individual := range population {
			fmt.Printf("%d -> %s\n", i, individual.RandomLetters)
		}
	}
}

func ToString(byteToConvert []byte) string {
	return strings.TrimSpace(string(byteToConvert))
}
