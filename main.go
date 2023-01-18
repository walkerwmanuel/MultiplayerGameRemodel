package main

import (
	"multiplayergame/myroutes"
	"multiplayergame/players"
)

func main() {

	err := players.ConnectDatabase()
	myroutes.CheckErr(err)

	myroutes.Routes()
}
