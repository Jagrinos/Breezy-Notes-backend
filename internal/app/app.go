package app

import (
	"BreeZy_Backend_vol_0/internal/db"
	"BreeZy_Backend_vol_0/internal/net"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() error {
	//start mongo + echo
	client, err := db.Connect()
	if err != nil {
		return err
	}

	e, err := net.SetupEcho()
	if err != nil {
		return err
	}
	go func() {
		if err = e.Start(":8008"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal(err)
		}
	}()

	//waiting for exit signal
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	//exit processing
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = e.Shutdown(ctx); err != nil {
		return err
	}
	if err = db.Disconnect(client); err != nil {
		return err
	}

	return nil
}

//func getAll(client *mongo.Client) {
//	users, err := users.GetAll(client)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if len(users) == 0 {
//		log.Println("no one")
//	}
//
//	for _, val := range users {
//		log.Printf("%s %s %s", val.Id, val.Login, val.Password)
//	}
//}
//var client *mongo.Client
//
//err := db.Connect(&client)
//if err != nil {
//return err
//}
//
//getAll(client)
//
//newuserid := uuid.NewString()
//newuser := Views.User{
//Id:       newuserid,
//Login:    "genadii_rabota",
//Password: "mypassword12",
//}
//users.Create(client, newuser)
//
//getAll(client)
//
//newuserChange := Views.UserWithoutId{
//Login:    "genadii",
//Password: "PIDORAS)",
//}
//
//users.Update(client, newuserid, newuserChange)
//
//getAll(client)
//
//users.Delete(client, newuserid)
//
//getAll(client)
//
//err = db.Disconnect(&client)
//if err != nil {
//return err
//}
