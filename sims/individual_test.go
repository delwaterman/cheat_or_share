package sims

import (
    "testing"
    "math/rand"
    "github.com/jacobsa/oglematchers"
    "github.com/jacobsa/ogletest"
)

// -------------------------------------------

type NewIndividualTest struct{}

func (t *NewIndividualTest) SetsMinAndMaxEnergy() {
    newIndividual := NewIndividual(1, 2)
    ogletest.ExpectThat((*newIndividual),
        oglematchers.DeepEquals(Individual{1, 2}))
}

func (t *NewIndividualTest) PanicsWhenMaxGreaterThanMin() {
    badIndividualFunc := func() {
        NewIndividual(2, 1)
    }
    expectedErrorMessage := "minEnergy(2) cannot be greater than maxEnergy(1)"
    ogletest.ExpectThat(badIndividualFunc,
                        oglematchers.Panics(
                            oglematchers.Equals(expectedErrorMessage)))
}

// -------------------------------------------

type NewRandIndividualTest struct {}

func (t *NewRandIndividualTest) IndividualMinimumAndMaximumIsCorrect() {
    var individual *Individual
    for count:= 0; count < 100; count ++ {
        individual = newRandIndividual()
        ogletest.ExpectThat(individual.MinEnergy, oglematchers.GreaterThan(0))
        ogletest.ExpectThat(individual.MaxEnergy, oglematchers.LessThan(11))
        ogletest.ExpectThat(individual.MinEnergy, oglematchers.LessOrEqual(individual.MaxEnergy))
    }
}

// -------------------------------------------

type IndividualExertEnergyTest struct {}

func (t *IndividualExertEnergyTest) ReturnsNumberBetweenMinAndMax() {
    var result int
    individual := Individual{1, 10}
    for i := 0; i < 100; i++ {
        result = individual.ExertEnergy()
        ogletest.ExpectThat(result, oglematchers.GreaterThan(0))
        ogletest.ExpectThat(result, oglematchers.LessThan(11))
    }
}

func (t *IndividualExertEnergyTest) ReturnsAtLeastMin() {
    individual := Individual{1, 10}
    // Seed 11 - rand.Int31n(1, 10) returns 0
    rand.Seed(11)
    ogletest.ExpectThat(individual.ExertEnergy(), oglematchers.Equals(1))
}


// -----------------------------------------

func init() {
    ogletest.RegisterTestSuite(&NewIndividualTest{})
    ogletest.RegisterTestSuite(&NewRandIndividualTest{})
    ogletest.RegisterTestSuite(&IndividualExertEnergyTest{})
}

func TestIndividual(t *testing.T) {
    ogletest.RunTests(t)
}

