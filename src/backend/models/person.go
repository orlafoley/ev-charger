package models

import (
	_ "github.com/mattn/go-sqlite3"
)

type Member struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func GetAllMember() ([]Member, error) {
	rows, err := DB.Query("SELECT id, username FROM MEMBERS")

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
