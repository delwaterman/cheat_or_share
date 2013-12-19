package sims

import (
    testing "testing"
    . "github.com/jacobsa/oglematchers"
    . "github.com/jacobsa/ogletest"
)

// -------------------------------------------

type NewWorldTest struct{
    cooperativeBuilder CooperativeBuilder
}

func (t *NewWorldTest) SetUp(testInfo *TestInfo) {
    individualBuilder := SimpleIndividualBuilder(1, 2)
    t.cooperativeBuilder = SimpleCooperativeBuilder(0, individualBuilder)
}

func (t *NewWorldTest) ReturnsWorldWithCooperatives() {
    world := newWorld(3, t.cooperativeBuilder)
    ExpectThat(len(world.Cooperatives), Equals(3))
}


// -----------------------------------------

func init() {
    RegisterTestSuite(&NewWorldTest{})
}

func TestWorld(t *testing.T) {
    RunTests(t)
}