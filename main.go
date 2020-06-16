package main

import (
	"fmt"
	"mavericks-consulting.com/assignment/rps/game"
)

func main() {
	fmt.Println("Welcome to RPS!!")
	gameMode, err := game.SelectGameMode()
	if err == nil {
		game.Play(gameMode)
	} else {
		fmt.Printf("%s\n", err.Error())
	}
}