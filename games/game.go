package games

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3" //Need to blank import package
)

var DB *sql.DB

// SqliteDB is the db object for sqlite database connections
var SqliteDB *sql.DB

func sqliteCreateTable(tableName string) error {

	statement, err := SqliteDB.Prepare("CREATE TABLE IF NOT EXISTS " + tableName + " (key TEXT NOT NULL UNIQUE PRIMARY KEY, value TEXT)")
	if err != nil {
		return err
	}

	defer statement.Close()
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	return nil
}

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func AddGame(newGame Game) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO games (players) VALUES (?)")

	if err != nil {
		fmt.Println("Error!")
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newGame.Players)

	if err != nil {
		fmt.Println("Error2!")
		return false, err
	}

	tx.Commit()

	return true, nil
}

func GetGame(count int) ([]Game, error) {

	//Simple seledt statement with a LIMIT appended to it
	rows, err := DB.Query("SELECT id, name, password, coins from people LIMIT" + strconv.Itoa(count))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	//Create a new slice people
	game := make([]Game, 0)

	for rows.Next() {
		//Create an instance of player struct
		singleGame := Game{}
		err = rows.Scan(&singleGame.Id, &singleGame.Players)

		if err != nil {
			return nil, err
		}

		game = append(game, singleGame)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return game, err
}
