package main

import (
	"sort"
)

func getIndividualFitness(i Individual) float64 {
	score := 0
	strToByte := []byte(EXPECTED_STR)
	for j := 0; j < LENGTH_OF_EXPECTED_STR && j < len(i.RandomLetters); j++ {
		if i.RandomLetters[j] == strToByte[j] {
			score++
		}
	}
	return float64(score)
}

func getAverageFitness(pop []Individual) float64 {
	averageScore := 0.0
	for j := 0; j < len(pop); j++ {
		averageScore += getIndividualFitness(pop[j])
	}
	return averageScore / POPULATION_COUNT
}

func gradePop(pop []Individual) []Individual {
	// returns a slice of individual in ascending order
	sort.Slice(pop, func(i, j int) bool {
		return getIndividualFitness(pop[i]) > getIndividualFitness(pop[j])
	})
	return pop
}

func goodResults(pop []Individual) []Individual {
	result := make([]Individual, 0)
	for i := 0; i < POPULATION_COUNT; i++ {
		if getIndividualFitness(pop[i]) == float64(MAXIMUM_FITNESS) {
			result = append(result, pop[i])
		}
	}
	return result
}

func getParents(pop []Individual) []Individual {
	popSortedWithScore := gradePop(pop)
	parents := make([]Individual, GRADED_INDIVIDUAL_RETAIN_COUNT)
	for i := 0; i < GRADED_INDIVIDUAL_RETAIN_COUNT; i++ {
		parents[i] = popSortedWithScore[i]
	}

	return parents
}
