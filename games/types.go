package games

import "multiplayergame/players"

type Game struct {
	Id      int
	Players []players.Player
}
