package main

import "fmt"

func createUser(user string) {
	fmt.Printf("user %s created\n", user)
}

func setPassword(user, password string) {
	fmt.Printf("user %s changed password to %s\n", user, password)
}

func createGroup(group string) {
	fmt.Printf("group %s created\n", group)
}

func createHomeDir(user string) {
	fmt.Printf("Directory /home/%s created\n", user)
}

func addUserFacade(user, password string) {
	createUser(user)
	setPassword(user, password)
	createGroup(user)
	createHomeDir(user)
}

func main() {
	addUserFacade("ernur", "134578")
}
