package suffixGenerator

import (
	"fmt"
	"math/rand"
)

// Generator takes the name from the user input and adds a random suffix to it
type Generator interface {
	Generate(name string) string
}

// SuffixGenerator is reponsible for randomly picking from the data
type SuffixGenerator struct {
	random *rand.Rand
}

// Generate creates a random suffix and returns the hero name
func (rn *SuffixGenerator) Generate(name string) string {
	randomAdjective := ADJECTIVES[rn.random.Intn(len(ADJECTIVES))]
	randomJob := ROLES[rn.random.Intn(len(ROLES))]

	randomName := fmt.Sprintf("%v, the %v %v", name, randomAdjective, randomJob)

	return randomName
}

// NewSuffixGenerator creates a new SuffixGenerator with a pseudo random number generator
func NewSuffixGenerator(seed int64) Generator {
	suffixGenerator := &SuffixGenerator{
		random: rand.New(rand.New(rand.NewSource(99))),
	}
	suffixGenerator.random.Seed(seed)

	return suffixGenerator
}
