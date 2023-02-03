package main

import (
	"multiplayergame/data"
	"multiplayergame/myroutes"
)

func main() {

	err := data.ConnectDatabase()
	myroutes.CheckErr(err)

	myroutes.Routes()
}
