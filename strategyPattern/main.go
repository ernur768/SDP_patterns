package main

import "fmt"

type Weapon interface {
	DealDamage() float32
}

type Sword struct {
	damage float32
}

func (s Sword) DealDamage() float32 {
	return s.damage
}

type GrenadeLauncher struct {
	damage float32
}

func (g GrenadeLauncher) DealDamage() float32 {
	return g.damage
}

type Player struct {
	weapon Weapon
}

func (p *Player) SetWeapon(weapon Weapon) {
	p.weapon = weapon
}

func main() {
	nagibator228 := Player{}

	nagibator228.SetWeapon(GrenadeLauncher{999.0})
	fmt.Println(nagibator228)

	nagibator228.SetWeapon(Sword{damage: 150.0})
	fmt.Println(nagibator228)
}
