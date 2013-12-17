package sims

import (
    testing "testing"
    . "github.com/jacobsa/oglematchers"
    . "github.com/jacobsa/ogletest"
)

// -------------------------------------------

type NewCooperativeWorkEffortTest struct{}

func (t *NewCooperativeWorkEffortTest) ReturnsEmptyCooperativeWorkEffort() {
    workEffort := newCooperativeWorkEffort()
    ExpectThat(len(workEffort.WorkByIndividual), Equals(0))
}

// -------------------------------------------

type TotalWorkTest struct {}

func (t *TotalWorkTest) ReturnsRandBuilderFunction() {
    workEffort := newCooperativeWorkEffort()
    workEffort.WorkByIndividual[&Individual{1, 2}] = 1
    workEffort.WorkByIndividual[&Individual{1, 2}] = 2
    workEffort.WorkByIndividual[&Individual{1, 2}] = 3

    ExpectThat(workEffort.TotalWork(), Equals(6))
}

// -----------------------------------------

func init() {
    RegisterTestSuite(&NewCooperativeWorkEffortTest{})
    RegisterTestSuite(&TotalWorkTest{})
}

func TestCooperativeWorkEffort(t *testing.T) {
    RunTests(t)
}