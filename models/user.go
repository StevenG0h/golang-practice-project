package models

import (
	"example.com/DB"
)

type User struct {
	Id       int
	Name     string
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) IsUserExists() bool {
	query := `
		SELECT COUNT(*) FROM users WHERE email = ?
	`
	row, err := DB.DB.Query(query, u.Email)

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

func (u User) GetUserPasswordByEmail() (string, error) {
	var retrievedPassword string
	query := `
		SELECT password FROM users WHERE email = ?
	`

	_, err := DB.DB.Prepare(query)

	if err != nil {
		panic(err)
	}

	err = DB.DB.QueryRow(query, u.Email).Scan(&retrievedPassword)

	if err != nil {
		panic(err)
	}

	if err != nil || retrievedPassword == "" {
		return "", err
	}

	return retrievedPassword, nil
}

func (u User) Save() error {
	query := `
		INSERT INTO users (name, email, password)
		VALUES (?, ?, ?)
	`
	_, err := DB.DB.Prepare(query)

	if err != nil {
		panic(err)
	}

	res, err := DB.DB.Exec(query, u.Name, u.Email, u.Password)

	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()

	if err != nil {
		panic(err)
	}

	u.Id = int(id)
	return err
}
