package main

import "fmt"

// Product interface ====================================

type BossBehavior interface {
	Attack()
}

// Concrete products ====================================

type Boss struct {
	behavior BossBehavior
}

// ------------------------------

type level1BossBehavior struct{}

func (l level1BossBehavior) Attack() {
	fmt.Println("harmless boss attacks")
}

type level2BossBehavior struct{}

func (l level2BossBehavior) Attack() {
	fmt.Println("ez attack pattern")
}

type level3BossBehavior struct{}

func (l level3BossBehavior) Attack() {
	fmt.Println("complex attack pattern")
}

// Factory ==============================================

func createBossBehavior(level string) BossBehavior {
	switch level {
	case "1":
		return level1BossBehavior{}
	case "2":
		return level2BossBehavior{}
	default:
		return level3BossBehavior{}
	}
}

// Client ===============================================

func main() {
	slimeBehavior := createBossBehavior("1")
	slime := Boss{behavior: slimeBehavior}
	slime.behavior.Attack()

	planterraBehavior := createBossBehavior("2")
	planterra := Boss{behavior: planterraBehavior}
	planterra.behavior.Attack()

	moonLordBehavior := createBossBehavior("3")
	moonLord := Boss{behavior: moonLordBehavior}
	moonLord.behavior.Attack()
}
