package sims

type CooperativeWorkEffort struct {
    WorkByIndividual map[*Individual]int
}

func newCooperativeWorkEffort() *CooperativeWorkEffort {
    workEffort := new(CooperativeWorkEffort)
    workEffort.WorkByIndividual = make(map[*Individual]int)
    return workEffort
}

func (workEffort *CooperativeWorkEffort) TotalWork() int {
    sum := 0
    for _, work := range workEffort.WorkByIndividual {
        sum += work
    }
    return sum
}