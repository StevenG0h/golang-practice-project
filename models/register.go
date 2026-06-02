package models

import "example.com/DB"

type Register struct {
	Id      int
	UserId  int64
	EventId int64
}

func (r Register) IsRegistrationExists() bool {
	query := `
		SELECT COUNT(*) FROM registers WHERE user_id = ? AND event_id = ?
	`
	row, err := DB.DB.Query(query, r.UserId, r.EventId)

	if err != nil {
		panic(err)
	}

	defer row.Close()

	var count int

	for row.Next() {
		err := row.Scan(&count)
		if err != nil {
			panic(err)
		}
	}

	return count > 0
}

func (r Register) Save() error {
	query := `
		INSERT INTO registers (user_id, event_id)
		VALUES (?, ?)
	`

	_, err := DB.DB.Prepare(query)

	if err != nil {
		panic(err)
	}

	res, err := DB.DB.Exec(query, r.UserId, r.EventId)

	if err != nil {
		panic(err)
	}

	_, err = res.LastInsertId()

	if err != nil {
		panic(err)
	}

	return err
}

func (r Register) Delete() error {
	query := `
		DELETE FROM registers WHERE user_id = ? AND event_id = ?
	`

	_, err := DB.DB.Prepare(query)

	if err != nil {
		panic(err)
	}

	_, err = DB.DB.Exec(query, r.UserId, r.EventId)

	if err != nil {
		panic(err)
	}

	return err
}
