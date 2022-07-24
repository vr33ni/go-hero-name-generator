package suffixgenerator

import (
	"testing"
	"time"
)

func TestHeroNameGenerator(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	_ = NewHeroNameGenerator(seed)
}

func TestHeroNameGenerator_GenerateWithEmptyName(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	heroNameGenerator := NewHeroNameGenerator(seed)

	name := heroNameGenerator.Generate("")
	if name == "" {
		t.Fatalf("Cannot create a hero name without a name.")
	}
}

func TestHeroNameGenerator_Generate(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	heroNameGenerator := NewHeroNameGenerator(seed)

	name := heroNameGenerator.Generate("Vreni")
	if name == "" {
		t.Fatalf("Cannot create a hero name without a name.")
	}
}
