package sims

import (

)

type Cooperative struct {
  Individuals []*Individual
}

func (cooperative *Cooperative) ExertEnergy() *CooperativeWorkEffort {
  exertedEnergy := newCooperativeWorkEffort()

  for _, individual := range cooperative.Individuals {
    exertedEnergy.WorkByIndividual[individual] = individual.ExertEnergy()
  }
  return exertedEnergy
}

func BuildNewCooperative(numIndividuals int,
    individualBuilder IndivdualBuilder) *Cooperative {

    cooperative := &Cooperative{}
    cooperative.Individuals = make([]*Individual, numIndividuals)
    for i := 0; i < numIndividuals; i++ {
        cooperative.Individuals[i] = individualBuilder()
    }

    return cooperative
}