package sims

type Individual struct {
    MinEnergy int
    MaxEnergy int
}

func NewIndividual(minEnergy int, maxEnergy int) Individual {
    individual := Individual{minEnergy, maxEnergy}
    return individual
}

func ExertEnergy(individual *Individual, randomNum float32) int {
    upperBound := individual.MaxEnergy - individual.MinEnergy + 1
    return int((randomNum * float32(upperBound)) + float32(individual.MinEnergy))
}