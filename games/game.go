package games

import (
	"fmt"
	"multiplayergame/data"
	"strconv"

	_ "github.com/mattn/go-sqlite3" //Need to blank import package
)

// SqliteDB is the db object for sqlite database connections

// func sqliteCreateTable(tableName string) error {

// 	statement, err := SqliteDB.Prepare("CREATE TABLE IF NOT EXISTS " + tableName + " (key TEXT NOT NULL UNIQUE PRIMARY KEY, value TEXT)")
// 	if err != nil {
// 		return err
// 	}

// 	defer statement.Close()
// 	_, err = statement.Exec()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// Function to open database file

func AddGame(newGame Game) (bool, error) {

	tx, err := data.DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO games (pot) VALUES (?)")

	if err != nil {
		fmt.Println("Error!")
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newGame.Pot)

	if err != nil {
		fmt.Println("Error2!")
		return false, err
	}

	tx.Commit()

	return true, nil
}

func AddPlayerToGame(gameid string, id string) (bool, error) {

	tx, err := data.DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO games (players) SELECT id FROM people WHERE id = ?")

	if err != nil {
		fmt.Println("Error!")
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(gameid)

	if err != nil {
		fmt.Println("Error2!")
		return false, err
	}

	tx.Commit()

	return true, nil
}

func GetGame(count int) ([]Game, error) {

	//Simple seledt statement with a LIMIT appended to it
	rows, err := data.DB.Query("SELECT id, pot, players from games LIMIT" + strconv.Itoa(count))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	//Create a new slice people
	game := make([]Game, 0)

	for rows.Next() {
		//Create an instance of player struct
		singleGame := Game{}
		err = rows.Scan(&singleGame.Id, &singleGame.Pot, &singleGame.Players)

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

func DeleteGame(gameId int) (bool, error) {

	fmt.Printf("Deleting player Id %d\n", gameId)

	tx, err := data.DB.Begin()

	if err != nil {
		return false, err
	}

	stmt, err := data.DB.Prepare("DELETE from games where id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(gameId)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}
