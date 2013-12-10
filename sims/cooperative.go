package sims

import "math/rand"

type Cooperative struct {
  Individuals []Individual
}

func (cooperative *Cooperative) ExertEnergy(randomSeed int64) int {
  rand.Seed(randomSeed)
  var exertedEnergy int
  for _, individual := range cooperative.Individuals {
    exertedEnergy += (&individual).ExertEnergy(rand.Float32())
  }
  return exertedEnergy
}