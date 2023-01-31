package main

import (
	"multiplayergame/myroutes"
)

func main() {

	err := myroutes.ConnectDatabase()
	myroutes.CheckErr(err)

	myroutes.Routes()
}
