package main

import (
    "github.com/delwaterman/cooperative_vs_competitive/sims"
    "fmt"
)

func main() {
    human := sims.NewIndividual(1, 10, 0)
    energy := sims.ExertEnergy(&human)
    fmt.Printf("Human %v\n", human)
    fmt.Printf("Did work at %v\n", energy)
}