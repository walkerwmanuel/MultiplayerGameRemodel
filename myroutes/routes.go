package myroutes

import (
	"fmt"
	"log"
	"multiplayergame/cardLogic"
	"multiplayergame/players"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()

	// API v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("person", getPersons)
		v1.GET("person/:id", getPersonById)
		v1.POST("person", addPerson)
		v1.PUT("person/:id", updatePerson)
		v1.DELETE("person/:id", deletePerson)
		v1.OPTIONS("person", options)
	}

	cLogic := r.Group("/api/cLogic")
	{
		cLogic.GET("shuffle", getShuffleDeck)
	}

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run()
}

func getPersons(c *gin.Context) {
	persons, err := players.GetPersons(10)
	CheckErr(err)

	if persons == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": persons})
	}
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getPersonById(c *gin.Context) {

	//Takes id param from web request and sets to a value
	id := c.Param("id")

	//Runs the GetPersonById function based on passed in web request id
	person, err := players.GetPersonById(id)
	CheckErr(err)
	// if the name is blank we can assume nothing is found
	if person.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": person})
	}
}

func addPerson(c *gin.Context) {

	//Creates instance of Player struct and names is json
	var json players.Player

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := players.AddPerson(json)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func updatePerson(c *gin.Context) {

	var json players.Player

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("First test")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	personId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Println("Second test")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	json.Id = personId
	success, err := players.UpdatePerson(json)

	if success {
		fmt.Println("3 test")
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		fmt.Println("4 test")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func deletePerson(c *gin.Context) {

	//Sets the id web input to "personId" variable
	personId, err := strconv.Atoi(c.Param("id"))

	//Checks for error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	//Calls DeletePerson function with respect to personId
	success, err := players.DeletePerson(personId)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func getShuffleDeck(c *gin.Context) {
	a := cardLogic.NewDeck()
	a.Shuffle()
	fmt.Printf("%v/n", a)
}

func options(c *gin.Context) {

	ourOptions := "HTTP/1.1 200 OK\n" +
		"Allow: GET,POST,PUT,DELETE,OPTIONS\n" +
		"Access-Control-Allow-Origin: http://locahost:8080\n" +
		"Access-Control-Allow-Methods: GET,POST,PUT,DELETE,OPTIONS\n" +
		"Access-Control-Allow-Headers: Content-Type\n"

	c.String(200, ourOptions)
}
