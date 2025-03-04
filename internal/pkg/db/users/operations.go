package users

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"uasbreezy/config/views"
)

func GetAll(d *sql.DB) ([]views.User, error) {
	var userls []views.User

	rows, err := d.Query(`SELECT * FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var us views.User
		err := rows.Scan(&us.Id, &us.Login, &us.Email, &us.About, &us.Password)
		if err != nil {
			log.Println(err)
			continue
		}
		userls = append(userls, us)
	}

	return userls, nil
}

func Create(d *sql.DB, u views.User) error {
	query := `
				INSERT INTO users (id, login, email, about, password)
				VALUES ($1, $2, $3, $4, $5)
			`

	hashedpass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = d.Exec(query, u.Id, u.Login, u.Email, u.About, hashedpass)
	if err != nil {
		return err
	}

	return nil
}

func Update(d *sql.DB, u views.UserNoId, id string) error {
	query := `
				UPDATE users
				SET login = $1, email = $2, about = $3, password = $4
				WHERE id = $5
			`

	hashedpass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = d.Exec(query, u.Login, u.Email, u.About, hashedpass, id)
	if err != nil {
		return err
	}
	return nil
}

func Delete(d *sql.DB, id string) error {
	query := `
				DELETE FROM users
				WHERE id = $1
			`

	_, err := d.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
