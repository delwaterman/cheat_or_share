package main

import (
    "github.com/delwaterman/cheat_or_share/sims"
    "fmt"
)

func main() {
    sims.InitRandomStrategy()

    randFunction := sims.GenericRandomStrategy()
    for call := 0; call < 100; call++ {
      fmt.Printf("Random call: %v, random number: %b\n", call, randFunction())
    }

    // human := sims.NewIndividual(1, 10, 0)
    // energy := sims.ExertEnergy(&human)
    // fmt.Printf("Human %v\n", human)
    // fmt.Printf("Did work at %v\n", energy)
}