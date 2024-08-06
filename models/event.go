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
	UserID      int64
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

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (E Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?,location = ?,dateTime = ?
	WHERE id = ?
	`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(E.Name, E.Description, E.Location, E.DateTime, E.Id)

	return err
}

func (E Event) DeleteEvent(id int64) error {
	query := `
	DELETE FROM events where id = ?
	`
	_, err := db.DB.Exec(query, id)
	return err

}

func (e *Event) Register(userId int64) error {
	query := `
	INSERT INTO registrations(event_id, user_id) VALUES (?,?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(e.Id, userId)
	return err
}

func (e *Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(e.Id, userId)
	return err
}
