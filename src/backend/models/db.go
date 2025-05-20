package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() error {
	var err error
	DB, err = sql.Open("sqlite3", "./ev_charger.db")
	if err != nil {
		log.Fatal(err)
	}

	// SQL Members Table Query
	createTableMembers := `DROP TABLE IF EXISTS MEMBERS;
		CREATE TABLE IF NOT EXISTS MEMBERS (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL
	);`

	_, err = DB.Exec(createTableMembers)
	if err != nil {
		log.Fatal(err)
	}

	insertSampleDataToMembers := `INSERT INTO MEMBERS(username)
									VALUES ("evan"),
											("orla"),
											("daniel");`
	_, err = DB.Exec(insertSampleDataToMembers)
	if err != nil {
		log.Fatal(err)
	}

	createTableBooking := `
	DROP TABLE IF EXISTS BOOKINGS;
	CREATE TABLE IF NOT EXISTS BOOKINGS (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			slot_id INTEGER NOT NULL,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			date TEXT NOT NULL,
			time TEXT NOT NULL,
			duration INTEGER NOT NULL
	);`
	_, err = DB.Exec(createTableBooking)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
