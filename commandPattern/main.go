package main

import "fmt"

// interface ==============================

type Command interface {
	execute()
}

type GameObject interface {
	goforward()
}

// concrete command =========================

type GoforwardCommand struct {
	gameObject GameObject
}

func (g GoforwardCommand) execute()  {   
	g.gameObject.goforward()
}

// invokers ==================================

type KeyboardButton struct {
	command Command
}

func (k KeyboardButton) press()  {   
	k.command.execute()
}

type GamepadButton struct {
	command Command
}

func (g GamepadButton) press()  {   
	g.command.execute()
}

// receiver ======================================

type Player struct {}

func (p Player) goforward()  {   
	fmt.Println("going forward")
}

// client ========================================

func main() {
	player := &Player{}
	
	command := GoforwardCommand{
		gameObject: player,
	}
	
	gamepadButton := GamepadButton{
		command: command,
	}
	
	keyboardButton := KeyboardButton{
		command: command,
	}
	
	gamepadButton.press()
	keyboardButton.press()
}