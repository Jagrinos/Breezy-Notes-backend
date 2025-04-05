package app

import (
	"BreeZy_Backend_vol_0/internal/net"
	"BreeZy_Backend_vol_0/internal/net/mongonet"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() error {
	//start mongo + echo
	var c mongonet.Client
	if err := c.Connect(); err != nil {
		return err
	}

	e, err := net.SetupEcho()
	if err != nil {
		return err
	}

	go func() {
		if err = e.Start(":8008"); err != nil && !errors.Is(http.ErrServerClosed, err) {
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
	if err = c.Disconnect(); err != nil {
		return err
	}

	return nil
}
