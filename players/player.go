package players

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

const PlayerTableName = "players"

var DB *sql.DB

// Function to open database file
func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

// My struct that uses "playerdata.db" file
type Player struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Coins    int    `json:"coins"`
	//Hand     cardLogic.Hand `json:"hand"`
}

func GetPersons(count int) ([]Player, error) {

	//Simple seledt statement with a LIMIT appended to it
	rows, err := DB.Query("SELECT id, name, password, coins from people LIMIT" + strconv.Itoa(count))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	//Create a new slice people
	people := make([]Player, 0)

	for rows.Next() {
		//Create an instance of player struct
		singlePerson := Player{}
		err = rows.Scan(&singlePerson.Id, &singlePerson.Name, &singlePerson.Password, &singlePerson.Coins)

		if err != nil {
			fmt.Println("Error here")
			return nil, err
		}

		people = append(people, singlePerson)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return people, err
}

// Function to find a specific player by their id
func GetPersonById(id string) (Player, error) {

	//Prepares for the QueryRow func
	//Important to prepare before QueryRow to check for error before sql func is called
	stmt, err := DB.Prepare("SELECT id, name, password, coins from people WHERE id = ?")

	if err != nil {
		return Player{}, err
	}

	//Create instance of player struct to be returned in function
	player := Player{}

	//After preparing for this earlier was can pass in the id and fine player who matches
	sqlErr := stmt.QueryRow(id).Scan(&player.Id, &player.Name, &player.Password, &player.Coins)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Player{}, nil
		}
		return Player{}, sqlErr
	}
	return player, nil
}

// Function to add to database
func AddPerson(newPerson Player) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO people (name, password, coins) VALUES (?, ?, ?)")

	if err != nil {
		fmt.Println("Error!")
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newPerson.Name, newPerson.Password, newPerson.Coins)

	if err != nil {
		fmt.Println("Error2!")
		return false, err
	}

	tx.Commit()

	return true, nil
}

func UpdatePerson(ourPlayer Player) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare(("UPDATE people SET name = ?, password = ?, coins = ? WHERE Id = ?"))

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(ourPlayer.Name, ourPlayer.Password, ourPlayer.Coins, ourPlayer.Id)

	if err != nil {
		return false, err
	}

	tx.Commit()
	return true, nil
}

func DeletePerson(playerId int) (bool, error) {

	fmt.Printf("Deleting player Id %d\n", playerId)

	tx, err := DB.Begin()

	if err != nil {
		return false, err
	}

	stmt, err := DB.Prepare("DELETE from people where id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(playerId)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}
