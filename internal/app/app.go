package app

import (
	"context"
	"errors"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"uasbreezy/config"
	"uasbreezy/internal/pkg/net"
)

func Run() error {
	err := config.SetupKeys()
	if err != nil {
		return err
	}
	driver, err := net.Connect()
	if err != nil {
		return err
	}

	e := net.GetEcho(driver)

	go func() {
		if err = e.Echo.Start(":8008"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
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
