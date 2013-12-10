package sims

import "testing"

func TestCooperativeExertEnergy(t *testing.T) {
  individual1 := NewIndividual(1, 10)
  individual2 := NewIndividual(1, 10)
  individual3 := NewIndividual(1, 10)

  cooperative := new(Cooperative)
  cooperative.Individuals = []Individual{individual1, individual2, individual3}

  // Check generation
  if output := (&cooperative).ExertEnergy(int64(0)); output != 20 {
    t.Errorf("(&cooperative).ExertEnergy() should return %v; got %v.", 20, output)
  }

  // Check that it generates the same value with same seed
  if output := (&cooperative).ExertEnergy(int64(0)); output != 20 {
    t.Errorf("(&cooperative).ExertEnergy() should return %v again; got %v.", 20, output)
  }
}