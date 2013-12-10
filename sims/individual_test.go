package sims

import "testing"

func TestIndividualExertEnergy(t *testing.T) {
  individual := NewIndividual(1, 10)

  // Test lower bound

  if output := (&individual).ExertEnergy(float32(0.0)); output != 1 {
    t.Errorf("ExertEnergy(individual) should return lower bound of %v; got %v.", 1, output)
  }


  if output := (&individual).ExertEnergy(float32(0.999999)); output != 10 {
    t.Errorf("ExertEnergy(individual) should return upper bound of %v; got %v.", 10, output)
  }
}