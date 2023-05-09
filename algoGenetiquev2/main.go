package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"

	_ "image/color"
)

// ----- Runtime configuration (edit at your convenience)

const (
	// Enter here the string to be searched
	expectedStr = "Fuck fucking fucked fucker fucking fuckups fuck fucking fucked fucking fuckup fucking fucker's fucking fuckup."

	// Enter here the chance for an individual to mutate (range 0-1)
	chanceToMutate = 0.1

	// Enter here the percent of top-grated individuals to be retained for the next generation (range 0-1)
	gradedRetainPercent = 0.2

	// Enter here the chance for a non top-grated individual to be retained for the next generation (range 0-1)
	chanceRetainNonGrated = 0.05

	// Number of individual in the population
	populationCount = 100

	// Maximum number of generation before stopping the script
	generationCountMax = 100000

	// ----- Do not touch anything after this line

	// Number of top-grated individuals to be retained for the next generation
	gradedIndividualRetainCount = int(float64(populationCount) * gradedRetainPercent)

	// Precompute the length of the expected string (individual are always fixed size objects)
	lengthOfExpectedStr = len(expectedStr)

	// Precompute LENGTH_OF_EXPECTED_STR // 2
	middleLengthOfExpectedStr = lengthOfExpectedStr / 2

	// Charmap of all allowed characters (A-Z a-z, space and !)
	allowedCharmap = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz !'."

	// Maximum fitness value
	maximumFitness = lengthOfExpectedStr
)

type Individual []rune

func getRandomChar() rune {
	// Return a random rune from the allowed charmap
	return rune(allowedCharmap[rand.Intn(len(allowedCharmap))])
}

func getRandomIndividual() []rune {
	// Create a new random individual
	individual := make([]rune, lengthOfExpectedStr)
	for i := range individual {
		individual[i] = getRandomChar()
	}
	return individual
}

func getRandomPopulation() []Individual {
	// Create a new random population, made of POPULATION_COUNT individual
	population := make([]Individual, populationCount)
	for i := range population {
		population[i] = Individual(getRandomIndividual())
	}
	return population
}

func getIndividualFitness(individual Individual) int {
	fitness := 0
	for i, c := range individual {
		if c == rune(expectedStr[i]) {
			fitness++
		}
	}
	return fitness
}

func averagePopulationGrade(population []Individual) float64 {
	total := 0
	for _, individual := range population {
		total += getIndividualFitness(individual)
	}
	return float64(total) / float64(populationCount)
}

func gradePopulation(population []Individual) []struct {
	individual Individual
	fitness    int
} {
	gradedIndividual := make([]struct {
		individual Individual
		fitness    int
	}, gradedIndividualRetainCount*populationCount)
	for i, individual := range population {
		gradedIndividual[i] = struct {
			individual Individual
			fitness    int
		}{individual, getIndividualFitness(individual)}
	}
	sort.Slice(gradedIndividual, func(i, j int) bool {
		return gradedIndividual[i].fitness > gradedIndividual[j].fitness
	})
	return gradedIndividual
}

func evolvePopulation(population []Individual) ([]Individual, float64, []Individual) {
	// Get individual sorted by grade (top first), the average grade and the solution (if any)
	rawGradedPopulation := gradePopulation(population)
	averageGrade := 0.0
	var solution []Individual
	gradedPopulation := make([]Individual, 0)
	for _, item := range rawGradedPopulation {
		individual, fitness := item.individual, item.fitness
		averageGrade += float64(fitness)
		gradedPopulation = append(gradedPopulation, individual)
		if fitness == maximumFitness {
			solution = append(solution, individual)
		}
	}
	averageGrade /= float64(populationCount)

	// End the script when solution is found
	if len(solution) > 0 {
		return population, averageGrade, solution
	}

	// Filter the top graded individuals
	parents := gradedPopulation[:gradedIndividualRetainCount]

	// Randomly add other individuals to promote genetic diversity
	for _, individual := range gradedPopulation[gradedIndividualRetainCount:] {
		if rand.Float64() < chanceRetainNonGrated {
			parents = append(parents, individual)
		}
	}

	// Mutate some individuals
	for _, individual := range parents {
		if rand.Float64() < chanceToMutate {
			placeToModify := rand.Intn(lengthOfExpectedStr)
			individual[placeToModify] = getRandomChar()
		}
	}

	// Crossover parents to create children
	parentsLen := len(parents)
	desiredLen := populationCount - parentsLen
	children := make([]Individual, 0, desiredLen)
	for len(children) < desiredLen {
		father := parents[rand.Intn(parentsLen)]
		mother := parents[rand.Intn(parentsLen)]
		if true { // father != mother
			child := append(father[:middleLengthOfExpectedStr], mother[middleLengthOfExpectedStr:]...)
			children = append(children, child)
		}
	}

	// The next generation is ready
	parents = append(parents, children...)
	return parents, averageGrade, solution
}

func main() {

	// Create a population and compute starting grade
	population := getRandomPopulation()
	average_grade := averagePopulationGrade(population)
	fmt.Printf("Starting grade: %.2f / %d\n", average_grade, maximumFitness)

	// Make the population evolve
	i := 0
	var solution []Individual
	log_avg := make([]float32, 0)
	for solution == nil && i < generationCountMax {
		population, average_grade, solution = evolvePopulation(population)
		if i&255 == 255 {
			fmt.Printf("Current grade: %.2f / %d (%d generation)\n", average_grade, maximumFitness, i)
		}
		if i&31 == 31 {
			log_avg = append(log_avg, float32(average_grade))
		}
		i++
	}

	// Print the final stats
	average_grade = averagePopulationGrade(population)
	print("Final grade: %.2f", average_grade, "/ %d\n", maximumFitness)

	// Print the solution
	if len(solution) > 1 {
		print("Solution found (%d times) after %d generations.\n", (len(solution)), i)
	} else {
		print("No solution found after %d generations.\n", i)
		print("- Last population was:\n")
		for number, individual := range population {
			strIndividual := make([]string, len(individual))
			for i, char := range individual {
				strIndividual[i] = string(char)
			}
			fmt.Printf("%d -> %s\n", number, strings.Join(strIndividual, ""))
		}
	}

}
