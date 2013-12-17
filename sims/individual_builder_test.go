package sims

import (
    "testing"
    "github.com/jacobsa/oglematchers"
    "github.com/jacobsa/ogletest"
)

type SimpleIndividualBuilderTest struct{}

func (t *SimpleIndividualBuilderTest) ReturnsSimpleIndiviualFunction() {
    var newIndividualFunction IndivdualBuilder
    newIndividualFunction = SimpleIndividualBuilder(1, 2)
    ogletest.ExpectThat((*newIndividualFunction()),
        oglematchers.DeepEquals(Individual{1, 2}))
}

// -------------------------------------------

type RandIndividualBuilderTest struct {}


func (t *RandIndividualBuilderTest) ReturnsRandBuilderFunction() {

    individualBuilderFunction := RandIndividualBuilder()
    ogletest.ExpectThat(individualBuilderFunction, oglematchers.Equals(newRandIndividual))
}

// -----------------------------------------

func init() {
    ogletest.RegisterTestSuite(&SimpleIndividualBuilderTest{})
    ogletest.RegisterTestSuite(&RandIndividualBuilderTest{})
}

func TestIndividualBuilder(t *testing.T) {
    ogletest.RunTests(t)
}