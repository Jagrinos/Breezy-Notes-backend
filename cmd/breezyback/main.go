package main

import (
	"BreeZy_Backend_vol_0/internal/app"
	"log"
)

func main() {
	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
