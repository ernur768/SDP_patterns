package main

import "fmt"

// Component interface ========================================

type Player interface {
	Play() string
}

// Concrete component ========================================

type NewPlayer struct {
}

func (n NewPlayer) Play() string {
	return "enjoying the game."
}

// Concrete decorators ========================================

// Decorator Tank --------------------

type Tank struct {
	Player Player
}

func (t Tank) Play() string {
	return t.Player.Play() + " Keeping the boss`s agro."
}

// Decorator DamageDealer

type DamageDealer struct {
	Player Player
}

func (d DamageDealer) Play() string {
	return d.Player.Play() + " Dealing high damage."
}

// Decorator Healer

type Healer struct {
	Player Player
}

func (h Healer) Play() string {
	return h.Player.Play() + " Healing teammates"
}

// client ========================================

func main() {

	noob := NewPlayer{}
	fmt.Printf("%s\n\n", noob.Play())

	pro := Tank{Player: noob}
	fmt.Printf("%s\n\n", pro.Play())

	hacker := DamageDealer{Player: pro}
	fmt.Printf("%s\n\n", hacker.Play())

	god := Healer{Player: hacker}
	fmt.Printf("%s\n\n", god.Play())
}
