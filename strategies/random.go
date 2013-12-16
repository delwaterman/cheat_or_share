package strategies

import (
    // "math/rand"
    // "time"
)

// type randStrategyFunction func() float32

type RandStrategy interface {
    RandomInteger(low, high int) int
}

// -----------------------------------

// type RealRandStrategy struct {
//     RandFunction randStrategyFunction
// }

// func (strategy *RealRandStrategy) RandomInteger(low, high int) int {
//     upperBound := high - low + 1
//     return int(strategy.RandFunction() * float32(upperBound)) + low
// }

// // Random Stratigies Functions
// func randomFloat32() float32 {
//    return rand.Float32()
// }


// func InitRandomStrategy() {
//     genericSeed := time.Now().UTC().UnixNano()
//     SetGenericRandSeed(genericSeed)
// }

// func SetGenericRandSeed(seed int64) {
//     rand.Seed(seed)
// }



