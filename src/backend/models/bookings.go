package models

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Bookings struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Date     string `json:"date"`
	Time     string `json:"time"`
	Duration int    `json:"duration"`
}

type BookingRequest struct {
	//SlotID   int    `json:"slot_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Date     string `json:"date" binding:"required"`
	Time     string `json:"time" binding:"required"`
	Duration int    `json:"duration" binding:"required"`
}

func InsertNewBooking(b Bookings) error {
	query := `INSERT INTO "BOOKINGS" ("name", "email", "date", "time", "duration")
              VALUES (?, ?, ?, ?, ?);`
	_, err := DB.Exec(query, b.Name, b.Email, b.Date, b.Time, b.Duration)
	if err != nil {
		log.Println("Error inserting booking into database:", err) // Log the error
	}
	return err
}

func GetAllBooking() ([]Bookings, error) {
	query := `SELECT * FROM "BOOKINGS"`
	bookings, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer bookings.Close()

	var bookingList []Bookings
	for bookings.Next() {
		var singleBooking Bookings
		err := bookings.Scan(&singleBooking.Id, &singleBooking.Name, &singleBooking.Email, &singleBooking.Date, &singleBooking.Time, &singleBooking.Duration)
		if err != nil {
			return nil, err
		}
		bookingList = append(bookingList, singleBooking)
	}
	return bookingList, nil
}
