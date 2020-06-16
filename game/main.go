package game

import (
	"fmt"
)

var playerVsComputer = ModeChoice{"Player Vs Computer"}
var computerVsComputer = ModeChoice{"Computer Vs Computer"}
var gameModeChoices = []ModeChoice{playerVsComputer, computerVsComputer}

var rock = Weapon{"Rock"}
var paper = Weapon{"Paper"}
var scissors = Weapon{"Scissors"}
var weapons = []Weapon{rock, paper, scissors}

var defeatRules = map[Weapon][]Weapon{
	rock:     []Weapon{scissors},
	paper:    []Weapon{rock},
	scissors: []Weapon{paper},
}

func Play(gameModeChoice ModeChoice) {
	switch gameModeChoice {
	case playerVsComputer:
		fmt.Println("Player Vs Computer is selected by the user")
		playPlayerVsComputer()
	case computerVsComputer:
		fmt.Println("Computer Vs Computer is selected by the user")
		playComputerVsComputer()
	}
}

func decideWinner(player1, player2 Player) Result {
	if player1.WeaponChoice == player2.WeaponChoice {
		return Result{
			IsDraw:       true,
			WinnerPlayer: player1,
			LoserPlayer:  player2}
	} else if contains(defeatRules[player1.WeaponChoice], player2.WeaponChoice) {
		return Result{
			IsDraw:       false,
			WinnerPlayer: player1,
			LoserPlayer:  player2}
	} else {
		return Result{
			IsDraw:       false,
			WinnerPlayer: player2,
			LoserPlayer:  player1}
	}
}
