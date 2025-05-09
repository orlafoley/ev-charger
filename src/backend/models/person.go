package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func ConnectDatabase() error {
	var err error
	db, err = sql.Open("sqlite3", "./ev_charger.db")
	if err != nil {
		log.Fatal(err)
	}

	// SQL Members Table Query
	createTableMembers := `DROP TABLE IF EXISTS MEMBERS;
		CREATE TABLE IF NOT EXISTS MEMBERS (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL
	);`

	_, err = db.Exec(createTableMembers)
	if err != nil {
		log.Fatal(err)
	}

	insertSampleDataToMembers := `INSERT INTO MEMBERS(username)
									VALUES ("evan"),
											("orla"),
											("daniel");`
	_, err = db.Exec(insertSampleDataToMembers)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

type Member struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func GetAllMember() ([]Member, error) {
	rows, err := db.Query("SELECT id, username FROM MEMBERS")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	member := make([]Member, 0)

	for rows.Next() {
		singleMember := Member{}
		err = rows.Scan(&singleMember.Id, &singleMember.Username)

		if err != nil {
			return nil, err
		}

		member = append(member, singleMember)
	}
	return member, err
}
