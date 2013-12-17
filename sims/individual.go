package sims

import (
    "fmt"
    "math/rand"
)

const (
    MinimumMinEnergy = 1
    MaximumMaxEnergy = 10
)

type Individual struct {
    MinEnergy int
    MaxEnergy int
}

func (individual *Individual) ExertEnergy() int {
    return randIntInclusive(individual.MinEnergy, individual.MaxEnergy)
}


func NewIndividual(minEnergy int, maxEnergy int) *Individual {
    if (maxEnergy < minEnergy) {
        panic(fmt.Sprintf("minEnergy(%v) cannot be greater than maxEnergy(%v)", minEnergy, maxEnergy))
    }
    return &Individual{minEnergy, maxEnergy}
}

func newRandIndividual() *Individual {
    lowerBound := randIntInclusive(MinimumMinEnergy, MaximumMaxEnergy)
    secondLowerBound := lowerBound
    upperBound := randIntInclusive(secondLowerBound, MaximumMaxEnergy)
    return NewIndividual(lowerBound, upperBound)
}

func randIntInclusive(minNum, maxNum int) int {
    upperBound := maxNum - minNum + 1
    return int(rand.Int31n(int32(upperBound))) + minNum
}

