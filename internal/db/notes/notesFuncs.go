package notes

import (
	"BreeZy_Backend_vol_0/config"
	"BreeZy_Backend_vol_0/models"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

//TODO USERS INDEX
/*
	{
		_id: ""
		title: ""
		create: time
		last_change: time
		users:[]
		//blocks: []
	}
*/

func Create(c *mongo.Client, n models.Note, ctx context.Context) error {
	if _, err := c.Database(config.Db).Collection(config.NotesColl).InsertOne(ctx, n); err != nil {
		return fmt.Errorf("NOTES: error in CREATE: %w", err)
	}
	return nil
}

func Delete(c *mongo.Client, id string, ctx context.Context) error {
	if _, err := c.Database(config.Db).Collection(config.NotesColl).DeleteOne(ctx, bson.D{{"_id", id}}); err != nil {
		return fmt.Errorf("NOTES: error in DELETE: %w", err)
	}
	return nil
}

func UpdateLastChange(c *mongo.Client, id string, ctx context.Context) error {
	if res, err := c.Database(config.Db).Collection(config.NotesColl).UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"last_change": time.Now().Local().Unix()}}); err != nil || res.MatchedCount == 0 {
		if res.MatchedCount == 0 {
			return fmt.Errorf("NOTES: error in UPDATE LAST CHANGE: %w", errors.New("not fiend note"))
		}
		return fmt.Errorf("NOTES: error in UPDATE LAST CHANGE: %w", err)
	}
	return nil
}

func GetAllByUser(c *mongo.Client, username string, ctx context.Context) ([]models.Note, error) {
	cur, err := c.Database(config.Db).Collection(config.NotesColl).Find(ctx, bson.M{"users": username})
	if err != nil {
		return nil, fmt.Errorf("NOTES: error in GET ALL BY USER: %w", err)
	}
	defer cur.Close(ctx)

	var notes []models.Note
	for cur.Next(ctx) {
		var n models.Note
		if err := cur.Decode(&n); err != nil {
			return notes, fmt.Errorf("NOTES: error in decode GET ALL BY USER: %w", err)
		}
		notes = append(notes, n)
	}

	return notes, nil
}
