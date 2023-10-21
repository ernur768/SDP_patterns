package main

import (
	"fmt"
	"sync"
)

var instanse *calculator
var mu sync.Mutex

func getCalculator() *calculator {
	if instanse == nil {
		mu.Lock()
		if instanse != nil {
			return instanse
		}
		instanse = &calculator{}
		mu.Unlock()
	}

	return instanse
}

type calculator struct{}

func (c calculator) sum(a, b int) {
	fmt.Println(a + b)
}

func main() {
	calc := getCalculator()

	calc.sum(5, 6)
}
