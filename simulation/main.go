package main

import (
    "math/rand"
    "github.com/delwaterman/cheat_or_share/sims"
    "fmt"
)

func main() {
    rand.Seed(0)

    cooperative := sims.BuildNewCooperative(5, sims.RandIndividualBuilder())
    fmt.Println("Cooperative:", cooperative)
    for ix, individual := range cooperative.Individuals {
        fmt.Printf(" --> Individual(%v): %v\n", ix, individual)
    }
    workEffort := cooperative.ExertEnergy()
    fmt.Println("ExertEnergy():")
    for individual, energy := range workEffort.WorkByIndividual {
        fmt.Printf(" --> Individual(%v): %v\n", individual, energy)
    }
    fmt.Printf("Total Energy: %v\n", workEffort.TotalWork())
}