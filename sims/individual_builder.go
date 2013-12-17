package sims

type IndivdualBuilder func() *Individual

func SimpleIndividualBuilder(minEnergy int, maxEnergy int) IndivdualBuilder {
    return func() *Individual {
        return NewIndividual(minEnergy, maxEnergy)
    }
}

func RandIndividualBuilder() IndivdualBuilder {
    return newRandIndividual
}