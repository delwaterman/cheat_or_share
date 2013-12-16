package sims

import (

)

type Cooperative struct {
  Individuals []*Individual
}

// func (cooperative *Cooperative) ExertEnergy(randomSeed int64) int {
//   rand.Seed(randomSeed)
//   var exertedEnergy int
//   for _, individual := range cooperative.Individuals {
//     exertedEnergy += (&individual).ExertEnergy(rand.Float32())
//   }
//   return exertedEnergy
// }

func BuildNewCooperative(numIndividuals int,
    individualBuilder IndivdualBuilder) *Cooperative {

    cooperative := &Cooperative{}
    cooperative.Individuals = make([]*Individual, numIndividuals)
    for i := 0; i < numIndividuals; i++ {
        cooperative.Individuals[i] = individualBuilder()
    }

    return cooperative
}