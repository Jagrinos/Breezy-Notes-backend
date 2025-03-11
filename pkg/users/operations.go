package users

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"uasbreezy/config/views"
)

func Auth(db views.SqlDb, u views.UserAuth) error {
	query := `
		SELECT login,password FROM users
		WHERE login = $1
	`

	var uf views.UserAuth
	err := db.QueryRow(query, u.Login).Scan(&uf.Login, &uf.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("no user found")
		}
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(uf.Password), []byte(u.Password))
	if err != nil {
		return errors.New("password incorrect")
	}

	return nil
}

func GetInfo(db views.SqlDb, login string) (views.UserInfo, error) {
	query := `
		SELECT login,email,about FROM users
		WHERE login = $1
	`
	var userinfo views.UserInfo

	err := db.QueryRow(query, login).Scan(&userinfo.Login, &userinfo.Email, &userinfo.About)
	if err != nil {
		return userinfo, err
	}

	return userinfo, nil
}
