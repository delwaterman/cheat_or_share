package sims

import "math/rand"

type Individual struct {
    MinEnergy int
    MaxEnergy int
    RandSeed int64
}

func NewIndividual(minEnergy int, maxEnergy int, randSeed int64) Individual {
    individual := Individual{minEnergy, maxEnergy, randSeed}

    return individual
}

func ExertEnergy(individual *Individual) int {
    upperBound := individual.MaxEnergy - individual.MinEnergy + 1
    rand.Seed(individual.RandSeed)
    return rand.Intn(upperBound) + individual.MinEnergy
}