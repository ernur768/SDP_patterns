package main

import (
	"fmt"
	"sync"
)

var instance *calculator
var mu sync.Mutex

func getCalculator() *calculator {
	if instance == nil {
		mu.Lock()
		if instance != nil {
			return instance
		}
		instance = &calculator{}
		mu.Unlock()
	}

	return instance
}

type calculator struct{}

func (c calculator) sum(a, b int) {
	fmt.Println(a + b)
}

func main() {
	calc := getCalculator()

	calc.sum(5, 6)
}
