package sims

import (
    "github.com/delwaterman/cheat_or_share/strategies"
    "fmt"
)

type Individual struct {
    MinEnergy int
    MaxEnergy int
}

type IndivdualBuilder func() *Individual


func NewIndividual(minEnergy int, maxEnergy int) *Individual {
    if (maxEnergy < minEnergy) {
        panic(fmt.Sprintf("minEnergy(%v) cannot be greater than maxEnergy(%v)", minEnergy, maxEnergy))
    }
    return &Individual{minEnergy, maxEnergy}
}

func SimpleIndividualBuilder(minEnergy int, maxEnergy int) IndivdualBuilder {
    return func() *Individual {
        return NewIndividual(minEnergy, maxEnergy)
    }
}

func NewRandIndividual(randomStrategy strategies.RandStrategy) *Individual {
    lowerBound := randomStrategy.RandomInteger(1, 10)

    secondLowerBound := lowerBound
    if lowerBound < 10 {
        secondLowerBound += 1
    }

    upperBound := randomStrategy.RandomInteger(secondLowerBound, 10)
    return NewIndividual(lowerBound, upperBound)
}

func RandIndividualBuilder(randomStrategy strategies.RandStrategy) IndivdualBuilder {
    return func() *Individual {
        return NewRandIndividual(randomStrategy)
    }
}

// func (individual *Individual) ExertEnergy(randomNum float32) int {
//     upperBound := (&individual).MaxEnergy - (&individual).MinEnergy + 1
//     return int((randomNum * float32(upperBound)) + float32(individual.MinEnergy))
// }