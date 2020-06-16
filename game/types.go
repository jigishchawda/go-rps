package game

type ModeChoice struct {
	name string
}

type Weapon struct {
	Name string
}

type Player struct {
	Name         string
	WeaponChoice Weapon
}

type Result struct {
	IsDraw       bool
	WinnerPlayer Player
	LoserPlayer  Player
}
