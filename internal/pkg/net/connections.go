package net

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
	"uasbreezy/config"
)

type DriverDb struct {
	Driver *sql.DB
}

func Connect() (DriverDb, error) {
	db, err := sql.Open("postgres", config.CONNSTR)
	if err != nil {
		return DriverDb{}, err
	}

	err = db.Ping()
	if err != nil {
		return DriverDb{}, err
	}

	log.Print("Connection is established")
	return DriverDb{Driver: db}, err
}

func Disconnect(d DriverDb) error {
	err := d.Driver.Close()
	if err != nil {
		return err
	}

	err = d.Driver.Ping()
	if err == nil {
		return errors.New("failed to disconnect")
	}
	log.Print("Connection terminated")
	return nil
}
