package users

import (
	"BreeZy_Backend_vol_0/Views"
	"BreeZy_Backend_vol_0/config"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

func Create(client *mongo.Client, user Views.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	coll := client.Database(config.Db).Collection(config.Usercoll)

	_, err := coll.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func Get(client *mongo.Client, id string) (error, Views.User) {
	coll := client.Database(config.Db).Collection(config.Usercoll)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user Views.User
	err := coll.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return err, user
	}
	return nil, user
}

func Update(client *mongo.Client, id string, updateData Views.UserWithoutId) error {
	coll := client.Database(config.Db).Collection(config.Usercoll)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := coll.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": updateData})
	if err != nil {
		return err
	}

	return nil
}

func Delete(client *mongo.Client, id string) error {
	coll := client.Database(config.Db).Collection(config.Usercoll)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := coll.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

func GetAll(client *mongo.Client) ([]Views.User, error) {
	uc := client.Database(config.Db).Collection(config.Usercoll)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := uc.Find(ctx, bson.D{})
	var users []Views.User

	if err != nil {
		return users, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var user Views.User
		if err := cur.Decode(&user); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
