package models

import (
	"errors"

	"project.com/event-booking/db"
	"project.com/event-booking/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {

	query := `INSERT INTO users(email,password) VALUES(?,?)`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	hashedPassord, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(u.Email, hashedPassord)
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id,password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.Id, &retrievedPassword)

	if err != nil {
		return err
	}
	pswdIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !pswdIsValid {
		return errors.New("cred invalid")
	}
	return nil
}
