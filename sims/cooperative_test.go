package sims

import (
    testing "testing"
    . "github.com/jacobsa/oglematchers"
    . "github.com/jacobsa/ogletest"
)

// -----------------------------------------

type BuildNewCooperativeTest struct {}

func (t *BuildNewCooperativeTest) ReturnsCooperative() {
    cooperative := BuildNewCooperative(3, SimpleIndividualBuilder(1, 2))
    individualSlices := cooperative.Individuals

    ExpectThat((*(individualSlices[0])),
        DeepEquals(Individual{1, 2}))
    ExpectThat((*(individualSlices[1])),
        DeepEquals(Individual{1, 2}))
    ExpectThat((*(individualSlices[2])),
        DeepEquals(Individual{1, 2}))
}

// -----------------------------------------

type CooperativeExertEnergyTest struct {
    Subject *Cooperative
    IndividualOneAddress *Individual
    IndividualTwoAddress *Individual
    IndividualThreeAddress *Individual
}

func (t *CooperativeExertEnergyTest) SetUp(testInfo *TestInfo) {
    t.IndividualOneAddress = &Individual{1, 1}
    t.IndividualTwoAddress = &Individual{2, 2}
    t.IndividualThreeAddress = &Individual{3, 3}

    count := 0
    individuals := []*Individual{t.IndividualOneAddress,
                                 t.IndividualTwoAddress,
                                 t.IndividualThreeAddress}
    privateBuilder := func() (individual *Individual) {
        individual = individuals[count]
        count ++
        return
    }

    t.Subject = BuildNewCooperative(3, privateBuilder)
}

func (t *CooperativeExertEnergyTest) ReturnsWorkEffort() {
    workEffort := t.Subject.ExertEnergy()

    ExpectThat(workEffort.WorkByIndividual[t.IndividualOneAddress],
        Equals(1))
    ExpectThat(workEffort.WorkByIndividual[t.IndividualTwoAddress],
        Equals(2))
    ExpectThat(workEffort.WorkByIndividual[t.IndividualThreeAddress],
        Equals(3))
}

// -----------------------------------------

func init() {
    RegisterTestSuite(&BuildNewCooperativeTest{})
    RegisterTestSuite(&CooperativeExertEnergyTest{})
}

func TestCooperative(t *testing.T) {
    RunTests(t)
}
