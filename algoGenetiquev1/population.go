package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
)

type Individual struct {
	RandomLetters []byte
}

type IndividualFitness struct {
	Individual Individual
	Fitness    int
}

func getRandomIndividual() *Individual {
	individual := new(Individual)
	individual.RandomLetters = make([]byte, LENGTH_OF_EXPECTED_STR)
	for i := 0; i < LENGTH_OF_EXPECTED_STR; i++ {
		individual.RandomLetters[i] = GetRandomChar()
	}
	return individual
}

func getRandomPop() []Individual {
	pop := make([]Individual, POPULATION_COUNT)
	for i := 0; i < POPULATION_COUNT; i++ {
		pop[i] = *getRandomIndividual()
	}
	return pop
}

func evolvePopulation(population []Individual) ([]Individual, float64, []Individual) {
	// Sort the population by fitness
	population = gradePop(population)

	// Create the next population with the best individuals
	nextPopulation := make([]Individual, len(population))

	// Replace the current population with the next population
	population = nextPopulation

	// Compute the average fitness of the population
	averageFitness := getAverageFitness(population)

	// Check if a solution has been found
	var solution []Individual
	for _, individual := range population {
		if int(getIndividualFitness(individual)) == MAXIMUM_FITNESS {
			solution = append(solution, individual)
		}
	}
	if len(solution) > 0 {
		return population, averageFitness, solution
	}

	return population, averageFitness, nil
}

func gradePopulation(population []Individual) ([]IndividualFitness, float64, []Individual) {
	var sumOfFitness float64
	var gradedPopulation []IndividualFitness
	var solution []Individual

	for _, individual := range population {
		fitness := getIndividualFitness(individual)
		gradedPopulation = append(gradedPopulation, IndividualFitness{individual, int(fitness)})
		sumOfFitness += fitness

		if fitness == float64(MAXIMUM_FITNESS) {
			solution = append(solution, individual)
		}
	}

	averageGrade := sumOfFitness / float64(len(population))

	sort.Slice(gradedPopulation, func(i, j int) bool {
		return gradedPopulation[i].Fitness > gradedPopulation[j].Fitness
	})

	return gradedPopulation, averageGrade, solution
}

func selectParents(gradedPopulation []IndividualFitness) []Individual {
	parents := make([]Individual, GRADED_INDIVIDUAL_RETAIN_COUNT)
	for i := range parents {
		parents[i] = gradedPopulation[i].Individual

	}
	return parents
}

func promoteDiversity(parents []Individual, gradedPopulation []IndividualFitness) []Individual {
	for _, individualFitness := range gradedPopulation[GRADED_INDIVIDUAL_RETAIN_COUNT:] {
		if rand.Float64() < CHANCE_RETAIN_NONGRATED {
			parents = append(parents, individualFitness.Individual)
		}
	}

	return parents
}

func mutateParents(parents []Individual) []Individual {
	for i := range parents {
		if rand.Float64() < CHANCE_TO_MUTATE {
			placeToModify := rand.Intn(len(parents[i].RandomLetters))
			parents[i].RandomLetters[placeToModify] = GetRandomChar()
		}
	}
	return parents
}

func crossoverParents(parents []Individual) []Individual {
	var children []Individual
	attempts := 0
	maxAttempts := 1000 // seuil maximal de tentatives

	for len(children) < POPULATION_COUNT-len(parents) {
		father := parents[rand.Intn(len(parents))]
		mother := parents[rand.Intn(len(parents))]

		if !bytes.Equal(father.RandomLetters, mother.RandomLetters) {
			childLetters := append(father.RandomLetters[:MIDDLE_LENGTH_OF_EXPECTED_STR], mother.RandomLetters[MIDDLE_LENGTH_OF_EXPECTED_STR:]...)
			child := Individual{childLetters}

			// vérifier si le nouveau enfant répond aux conditions pour être ajouté
			if getIndividualFitness(child) > 0 {
				children = append(children, child)
				attempts = 0 // réinitialiser le compteur
			}
		}

		attempts++

		// si on a dépassé le seuil maximal de tentatives, sortir de la boucle
		if attempts > maxAttempts {
			fmt.Println("maximum attempts reached, exiting the loop")
			break
		}
	}

	return children
}
