package sims

import (
    testing "testing"
    . "github.com/jacobsa/oglematchers"
    . "github.com/jacobsa/ogletest"
)

type SimpleCooperativeBuilderTest struct{}

func (t *SimpleCooperativeBuilderTest) ReturnsSimpleCooperativeFunction() {
    var newCooperativeBuilderFunc CooperativeBuilder
    newCooperativeBuilderFunc = SimpleCooperativeBuilder(3, SimpleIndividualBuilder(1, 2))
    cooperative := newCooperativeBuilderFunc()
    ExpectThat(len(cooperative.Individuals), Equals(3))
}

// -----------------------------------------

func init() {
    RegisterTestSuite(&SimpleCooperativeBuilderTest{})
}

func TestCooperativeBuilder(t *testing.T) {
    RunTests(t)
}