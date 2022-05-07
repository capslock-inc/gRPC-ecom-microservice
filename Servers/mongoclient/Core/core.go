package core

import (
	"context"
	"log"

	model "github.com/capslock-inc/gprc-demo/Servers/mongoclient/Model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// create cart
func CreateCart(data *model.CartItem, client *mongo.Client) (string, error) {
	collection := client.Database("cartdb").Collection("cart")

	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatalf("error adding cartitem : %v", err)
		return string(""), err
	} else {
		return string("cart created sucessfully"), nil
	}
}

// for testing purpose (will return all documents)
func GetCart(client *mongo.Client) []model.CartItem {
	collection := client.Database("cartdb").Collection("cart")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatalf("error getting item: %v ", err)
	}
	var Allcart []model.CartItem
	for cursor.Next(context.TODO()) {
		var item model.CartItem
		cursor.Decode(&item)
		log.Println(item)
		Allcart = append(Allcart, item)
	}
	defer cursor.Close(context.TODO())
	return Allcart
}
