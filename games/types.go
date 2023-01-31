package games

import "multiplayergame/players"

type Game struct {
	Id      int
	Pot     int
	Players []players.Player
}
