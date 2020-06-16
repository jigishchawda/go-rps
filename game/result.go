package game

import(
	"fmt"
)

func (result Result) Display() {
	fmt.Printf("%s player selected: %s\n",
		result.WinnerPlayer.Name,
		result.WinnerPlayer.WeaponChoice.Name)

	fmt.Printf("%s player selected: %s\n",
		result.LoserPlayer.Name,
		result.LoserPlayer.WeaponChoice.Name)

	if result.IsDraw {
		fmt.Println("Its a Draw!!!")
	} else {
		fmt.Printf("%s player wins!!!\n", result.WinnerPlayer.Name)
	}
}