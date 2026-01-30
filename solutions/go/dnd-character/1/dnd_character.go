package dndcharacter

import (
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier by subtracting 10 and
// performing an arithmetic right shift to achieve floor(diff / 2).
func Modifier(score int) int {
	return (score - 10) >> 1
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	total := 0
	min := 7
	for range 4 {
		tmp := rand.Intn(6) + 1
		if tmp < min {
			min = tmp
		}
		total += tmp
	}
	return total - min
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	constitution := Ability()
	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: constitution,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(constitution),
	}
}
