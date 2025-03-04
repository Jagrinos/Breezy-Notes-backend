package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"uasbreezy/internal/pkg/net"
)

func Run() error {
	driver, err := net.Connect()
	if err != nil {
		return err
	}

	//usersss, err := users.GetAll(driver.Driver)
	//log.Print(usersss)

	//err = users.Create(driver, views.User{
	//	Id:       uuid.NewString(),
	//	Login:    "1",
	//	Email:    "1",
	//	Password: "1",
	//	About:    "1",
	//})
	//usersss, err := users.GetAll(driver)
	//log.Print(usersss)
	//
	//if err != nil {
	//	return err
	//}
	//usersls, err := GetAll(driver)
	//if err != nil {
	//	if err = net.Disconnect(driver); err != nil {
	//		return err
	//	}
	//	return err
	//}
	//
	//log.Print(usersls)

	e := net.GetEcho(driver)

	go func() {
		if err = e.Echo.Start(":8008"); err != nil && err != http.ErrServerClosed {
			log.Print(err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	if err = e.Echo.Shutdown(context.Background()); err != nil {
		if err = net.Disconnect(driver); err != nil {
			return err
		}
		return err
	}

	if err = net.Disconnect(driver); err != nil {
		return err
	}

	return nil
}
