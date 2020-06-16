package game

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func SelectGameMode() (ModeChoice, error) {
	fmt.Println("Select Game Mode to play")
	for index, choice := range gameModeChoices {
		fmt.Printf("%d %s\n", index+1, choice.name)
	}
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	choice, err := strconv.Atoi(text)
	if err == nil && choice > 0 && choice <= len(gameModeChoices) {
		return gameModeChoices[choice-1], nil
	}
	return ModeChoice{}, errors.New("incorrect selection")
}

func playPlayerVsComputer() {
	computerPlayer := Player{"Computer", randomWeapon()}

	playerWeapon, err := selectWeapon()
	if err == nil {
		humanPlayer := Player{"Human", playerWeapon}
		gameResult := decideWinner(humanPlayer, computerPlayer)
		gameResult.Display()
	} else {
		fmt.Println(err.Error())
	}
}

func playComputerVsComputer() {
	computerPlayer1 := Player{"Computer 1", randomWeapon()}
	computerPlayer2 := Player{"Computer 2", randomWeapon()}
	gameResult := decideWinner(computerPlayer1, computerPlayer2)
	gameResult.Display()
}

func selectWeapon() (Weapon, error) {
	fmt.Println("Select a weapon")
	for index, weaponOption := range weapons {
		fmt.Printf("%d. %s\n", index+1, weaponOption.Name)
	}
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	choice, err := strconv.Atoi(text)
	if err == nil && choice > 0 && choice <= len(weapons) {
		return weapons[choice-1], nil
	}
	return Weapon{}, errors.New("Wrong selection")
}

func randomWeapon() Weapon {
	value := time.Now().UnixNano()
	source := rand.NewSource(value)
	random := rand.New(source)
	randomIndex := random.Intn(len(weapons))
	return weapons[randomIndex]
}
