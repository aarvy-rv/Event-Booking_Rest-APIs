package models

import (
	"time"

	"project.com/event-booking/db"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events []Event = []Event{}

func (e *Event) Save() error {

	query := `INSERT INTO events(name,description,location,dateTime,user_id)
	VALUES(?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)
	//When we prepare a memory we store it in memory,to reuse further
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}
	e.Id, err = result.LastInsertId()
	if err != nil {
		return err
	}
	events = append(events, *e)
	return nil

}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query) //Using query instead of exec as query is used when rows are returned whereas Exec() is used
	// when data is updated in the table i.e insert/update/delete

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var eventData []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		eventData = append(eventData, event)

	}
	return eventData, nil
}
