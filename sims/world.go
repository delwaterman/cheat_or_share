package sims

type World struct {
    Cooperatives []*Cooperative
}

func newWorld(numCooperatives int, cooperativeBuilder CooperativeBuilder) (world *World) {
    world = new(World)
    world.Cooperatives = make([]*Cooperative, numCooperatives)
    for ix := 0; ix < numCooperatives; ix++ {
        world.Cooperatives[ix] = cooperativeBuilder()
    }

    return
}