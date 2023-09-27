package main

import "fmt"

// interfaces ========================================
type Weapon interface {
	deal_damage() float32
}

type Armor interface {
	defend(float32) float32
}

// structs and methods ========================================

// Weapon implementation ---------------

type Knife struct {
	damage float32
}

func (k Knife) deal_damage() float32 {
	return k.damage
}

type Grenade_launcher struct {
	damage float32
}

func (g Grenade_launcher) deal_damage() float32 {
	return g.damage
}

// Armor implementation -------------

type Body_armor struct {
	endurance float32
}

func (b Body_armor) defend(damage float32) float32 {
	finalDamage := damage - b.endurance
	if finalDamage <= 0 {
		return 1.0
	} else {
		return finalDamage
	}
}

// Context ========================================

type Player struct {
	Weapon1 Weapon
	Weapon2 Weapon
	Armor   Armor
}

// Client ==================================================

func main() {

	AGS_40_Balkan := Grenade_launcher{damage: 999.0}
	katana := Knife{damage: 150.0}
	kevlar_armor := Body_armor{endurance: 60.0}

	nagibator228 := Player{
		Weapon1: AGS_40_Balkan,
		Weapon2: katana,
		Armor:   kevlar_armor}

	fmt.Println(nagibator228)
}
