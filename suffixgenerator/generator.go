package suffixgenerator

import (
	"fmt"
	"math/rand"
)

// Generator takes the name from the user input and adds a random suffix to it
type HeroNameGenerator interface {
	Generate(name string) string
}

// HeroNameGenerator is reponsible for randomly picking from the data
type heroNameGenerator struct {
	random *rand.Rand
}

// Generate creates a random suffix and returns the hero name
func (rn *heroNameGenerator) Generate(name string) string {
	randomAdjective := ADJECTIVES[rn.random.Intn(len(ADJECTIVES))]
	randomJob := ROLES[rn.random.Intn(len(ROLES))]

	randomName := fmt.Sprintf("%v, the %v %v", name, randomAdjective, randomJob)

	return randomName
}

// NewHeroNameGenerator creates a new SuffixGenerator with a pseudo random number generator
func NewHeroNameGenerator(seed int64) HeroNameGenerator {
	heroNameGenerator := &heroNameGenerator{
		random: rand.New(rand.New(rand.NewSource(99))),
	}
	heroNameGenerator.random.Seed(seed)

	return heroNameGenerator
}
