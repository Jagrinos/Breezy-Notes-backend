package mongonet

import (
	"BreeZy_Backend_vol_0/config"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"time"
)

type Client struct {
	Client *mongo.Client
}

func (c *Client) Connect() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c.Client, err = mongo.Connect(options.Client().ApplyURI(config.Uri))
	if err != nil {
		return fmt.Errorf("error in connection to mongo: %w", err)
	}
	if err = c.Client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("PING: connection not established: %w", err)
	}

	log.Println("Connection to MongoDB is established successfully")

	return nil
}

func (c *Client) Disconnect() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = c.Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("error when terminating connection to mongo: %w", err)
	}
	if err = c.Client.Ping(ctx, nil); err == nil {
		return errors.New("PING: terminating connection to MongoDB is failed")
	}
	log.Println("Connection to MongoDB is terminated successfully")

	return nil
}
