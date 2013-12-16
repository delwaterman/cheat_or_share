package sims

import (
    "testing"
    "github.com/jacobsa/oglematchers"
    "github.com/jacobsa/ogletest"
    "github.com/delwaterman/cheat_or_share/mock_strategies"
)

// -----------------------------------------

type BuildNewCooperativeTest struct {
    RandStrategyMock mock_strategies.MockRandStrategy
}



func (t *BuildNewCooperativeTest) ReturnsCooperative() {
    cooperative := BuildNewCooperative(3, SimpleIndividualBuilder(1, 2))
    individualSlices := cooperative.Individuals

    ogletest.ExpectThat((*(individualSlices[0])),
        oglematchers.DeepEquals(Individual{1, 2}))
    ogletest.ExpectThat((*(individualSlices[1])),
        oglematchers.DeepEquals(Individual{1, 2}))
    ogletest.ExpectThat((*(individualSlices[2])),
        oglematchers.DeepEquals(Individual{1, 2}))
}


// -----------------------------------------

func init() {
    ogletest.RegisterTestSuite(&BuildNewCooperativeTest{})
}

func TestCooperative(t *testing.T) {
    ogletest.RunTests(t)
}


// func TestCooperativeExertEnergy(t *testing.T) {
//   individual1 := NewIndividual(1, 10)
//   individual2 := NewIndividual(1, 10)
//   individual3 := NewIndividual(1, 10)

//   cooperative := new(Cooperative)
//   cooperative.Individuals = []Individual{individual1, individual2, individual3}

//   // Check generation
//   if output := (&cooperative).ExertEnergy(int64(0)); output != 20 {
//     t.Errorf("(&cooperative).ExertEnergy() should return %v; got %v.", 20, output)
//   }

//   // Check that it generates the same value with same seed
//   if output := (&cooperative).ExertEnergy(int64(0)); output != 20 {
//     t.Errorf("(&cooperative).ExertEnergy() should return %v again; got %v.", 20, output)
//   }
// }