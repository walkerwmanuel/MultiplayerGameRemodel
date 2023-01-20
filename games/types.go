package games

import "multiplayergame/players"

type Game struct {
	GameId  int
	Players []players.Player
}
