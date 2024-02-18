package template

import (
	"math/rand"
	"time"
)

type Randomizer struct {
	Rand *rand.Rand
}

// NewRandomizer creates a new Randomizer with a default seed based on the current Unix timestamp.
func NewRandomizer() *Randomizer {
	seed := time.Now().UnixNano()
	return &Randomizer{
		Rand: rand.New(rand.NewSource(seed)),
	}
}

// NewRandomizerWithSeed creates a new Randomizer with the provided seed.
func NewRandomizerWithSeed(seed int64) *Randomizer {
	return &Randomizer{
		Rand: rand.New(rand.NewSource(seed)),
	}
}
