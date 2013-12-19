package sims

type CooperativeBuilder func() *Cooperative

func SimpleCooperativeBuilder(numIndividuals int,
                              individualBuilder IndivdualBuilder) CooperativeBuilder {
    return func() *Cooperative {
        return BuildNewCooperative(numIndividuals, individualBuilder)
    }
}