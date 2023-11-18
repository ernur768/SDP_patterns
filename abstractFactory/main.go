package main

import "fmt"

type Button interface {
	Paint()
}

type Checkbox interface {
	Paint()
}

type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

type WinButton struct{}

func (wb WinButton) Paint() {
	fmt.Println("Windows Button")
}

type WinCheckbox struct{}

func (wc WinCheckbox) Paint() {
	fmt.Println("Windows Checkbox")
}

type WinFactory struct{}

func (w WinFactory) CreateButton() Button { return WinButton{} }

func (w WinFactory) CreateCheckbox() Checkbox { return WinCheckbox{} }

func main() {
	factory := WinFactory{}
	factory.CreateButton().Paint()
	factory.CreateCheckbox().Paint()
}
