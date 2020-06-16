package game

func contains(weapons []Weapon, w Weapon) bool {
	for _, item := range weapons {
		if item == w {
			return true
		}
	}
	return false
}