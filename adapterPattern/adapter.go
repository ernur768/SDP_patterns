package main

import "fmt"

type Probability interface {
	GetProbability() float64
}

type Calculator struct{}

func (Calculator) GetProbability() float64 {
	return 20.0 / 100
}

type CalculatorAdapter struct {
	calculator Calculator
}

func (ca CalculatorAdapter) GetProbability() float64 {
	return ca.calculator.GetProbability() * 100
}

type Screen struct {
	p Probability
}

func (s Screen) show() {
	fmt.Printf("Probability is %v percents", s.p.GetProbability())
}

func main() {
	calculator := Calculator{}
	adapter := CalculatorAdapter{calculator}
	screen := Screen{adapter}
	screen.show()
}
