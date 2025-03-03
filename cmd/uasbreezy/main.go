package main

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"uasbreezy/config"
	"uasbreezy/config/views"
)

func main() {
	//connection

	db, err := sql.Open("postgres", config.CONNSTR)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("success connect")

	//insert
	newid := uuid.NewString()
	u := views.User{
		Id:       newid,
		Login:    "provorov2",
		Email:    "provorov@admin.com",
		Password: "provorovprovorov",
		About:    "my provorovprovorovprovorov im provorovprovorovprovorovprovorov we are provorovprovorovprovorov",
	}

	query := `
		INSERT INTO users (id, login, email, password, about)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	err = db.QueryRow(query, u.Id, u.Login, u.Email, u.Password, u.About).Scan(&u.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("User created with ID:", u.Id)

	rows, err := db.Query(`SELECT * FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var u views.User
		err := rows.Scan(&u.Id, &u.Login, &u.Email, &u.Password, &u.About)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Print(u)
	}
}
