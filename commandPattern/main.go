package main

import "fmt"

type Command interface {
	execute()
}

type Player struct{}

func (p Player) goForward() {
	fmt.Println("going forward")
}

type GoForwardCommand struct {
	Player
}

func (gfc GoForwardCommand) execute() {
	gfc.Player.goForward()
}

type KeyboardButton struct {
	command Command
}

func (kb KeyboardButton) press() {
	kb.command.execute()
}

type GamepadButton struct {
	command Command
}

func (gb GamepadButton) press() {
	gb.command.execute()
}

func main() {
	player := Player{}
	command := GoForwardCommand{player}

	gamepadButton := GamepadButton{command: command}
	keyboardButton := KeyboardButton{command: command}

	gamepadButton.press()
	keyboardButton.press()
}
