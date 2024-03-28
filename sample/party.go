package main

import "github.com/kharism/tbcombat"

var (
	Warrior1 *tbcombat.Character
)

func NewWarrior() *tbcombat.Character {
	warrior := &tbcombat.Character{Name: "Warrior1", HP: 20, MaxHP: 20}

	return warrior
}
