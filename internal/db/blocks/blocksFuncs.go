package blocks

import (
	"BreeZy_Backend_vol_0/config"
	"BreeZy_Backend_vol_0/internal/db/notes"
	"BreeZy_Backend_vol_0/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func Create(c *mongo.Client, b models.Block, ctx context.Context) error {
	if _, err := c.Database(config.Db).Collection(config.BlockColl).InsertOne(ctx, b); err != nil {
		return fmt.Errorf("BLOCKS: error in CREATE: %w", err)
	}

	return nil
}

func Delete(c *mongo.Client, id string, ctx context.Context) error {
	coll := c.Database(config.Db).Collection(config.BlockColl)
	if _, err := coll.DeleteOne(ctx, bson.D{{"_id", id}}); err != nil {
		return fmt.Errorf("BLOCKS: error in DELETE: %w", err)
	}

	return nil
}

// TODO TEST 'CAUSE VERY OFTEN USE
func UpdateContent(c *mongo.Client, id string, updateData []models.TextBlock, ctx context.Context) error {
	exit := make(chan error, 1)
	go func() {
		cur := c.Database(config.Db).Collection(config.BlockColl).FindOne(ctx, bson.M{"_id": id})
		var b models.Block
		if err := cur.Decode(&b); err != nil {
			exit <- err
			return
		}

		if err := notes.UpdateLastChange(c, b.NoteId, ctx); err != nil {
			exit <- err
			return
		}
		exit <- nil
	}()

	if _, err := c.Database(config.Db).Collection(config.BlockColl).UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"content": updateData}}); err != nil {
		return fmt.Errorf("BLOCKS: error in UPDATE CONTENT: %w", err)
	}
	if err := <-exit; err != nil {
		return fmt.Errorf("BLOCKS: error in UPDATE CONTENT bad note ID: %w", err)
	}
	return nil
}

func GetAllInNote(c *mongo.Client, noteId string, ctx context.Context) ([]models.Block, error) {
	cur, err := c.Database(config.Db).Collection(config.BlockColl).Find(ctx, bson.M{"noteId": noteId})
	if err != nil {
		return nil, fmt.Errorf("BLOCKS: error in GET ALL IN NOTE: %w", err)
	}
	defer cur.Close(ctx)

	var blocks []models.Block
	for cur.Next(ctx) {
		var block models.Block
		if err := cur.Decode(&block); err != nil {
			return blocks, fmt.Errorf("BLOCKS: error in decode GET ALL IN NOTE: %w", err)
		}
		blocks = append(blocks, block)
	}

	return blocks, nil
}
