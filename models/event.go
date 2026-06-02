package models

import (
	"time"

	"example.com/DB"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Datetime    time.Time `binding:"required"`
	UserID      int
}

func (e *Event) Save() error {
	query := `
		INSERT INTO events (name, description, location, datetime, user_id)
		VALUES (?, ?, ?, ?, ?)
	`
	_, err := DB.DB.Prepare(query)

	if err != nil {
		panic(err)
	}

	res, err := DB.DB.Exec(query, e.Name, e.Description, e.Location, e.Datetime, e.UserID)

	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()

	if err != nil {
		panic(err)
	}

	e.ID = int(id)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `
		SELECT * FROM events
	`

	rows, err := DB.DB.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetById(id int) (Event, error) {
	query := `
		SELECT * FROM events WHERE id = ?
	`
	row, err := DB.DB.Query(query, id)

	if err != nil {
		return Event{}, err
	}

	defer row.Close()

	var event Event

	for row.Next() {
		err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)
		if err != nil {
			return Event{}, err
		}
	}

	return event, nil
}

func (event *Event) UpdateById() error {
	query := `
		UPDATE events SET name = ?, description = ?, location = ?, datetime = ?, user_id = ? WHERE id = ?
	`
	_, err := DB.DB.Prepare(query)

	if err != nil {
		panic(err)
	}

	res, err := DB.DB.Exec(query, event.Name, event.Description, event.Location, event.Datetime, event.UserID, event.ID)

	if err != nil {
		panic(err)
	}
	rows, err := res.RowsAffected()

	if err != nil && rows == 0 {
		panic(err)
	}

	return err
}

func DeleteById(id int) error {
	query := `
		DELETE FROM events WHERE id = ?
	`

	_, err := DB.DB.Prepare(query)

	if err != nil {
		return err
	}

	_, err = DB.DB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
