package main

import "fmt"

// Interfaces ========================

type LeftShoe interface {
	IsOnLeftLeg() bool
}

type RightShoe interface {
	IsOnRightLeg() bool
}

type ShoeFactory interface {
	createLeftShoe() LeftShoe
	createRightShoe() RightShoe
}

// Concrete products =============================

type ClownLeftShoe struct{}

func (c ClownLeftShoe) IsOnLeftLeg() bool {
	return true
}

type ClownRightShoe struct{}

func (c ClownRightShoe) IsOnRightLeg() bool {
	return true
}

// Concrete factory

type ClownShoeFactory struct{}

func (c ClownShoeFactory) createLeftShoe() LeftShoe {
	return &ClownLeftShoe{}
}

func (c ClownShoeFactory) createRightShoe() RightShoe {
	return &ClownRightShoe{}
}

// Client ==================================

func main() {
	factory := ClownShoeFactory{}
	LeftShoe := factory.createLeftShoe()
	RightShoe := factory.createRightShoe()
	fmt.Println(LeftShoe, RightShoe)
}
