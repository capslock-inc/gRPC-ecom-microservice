package core

import (
	"context"
	"log"

	model "github.com/capslock-inc/gprc-demo/Servers/mongoclient/Model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetCartItemByUserId
func GetCartItemByUserId(userid string, client *mongo.Client) (model.CartItem, error) {
	collection := client.Database("cartdb").Collection("cart")
	var cart model.CartItem
	err := collection.FindOne(context.TODO(), bson.D{{"_id", userid}}).Decode(&cart)
	if err != nil {
		log.Fatalf("err getting cart : %v", err)
		return model.CartItem{}, err
	}
	return cart, nil
}

// CreateNewCart
func CreateNewCart(data *model.CartItem, client *mongo.Client) (string, error) {
	collection := client.Database("cartdb").Collection("cart")

	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatalf("error adding cartitem : %v", err)
		return string(""), err
	} else {
		return string("cart created sucessfully"), nil
	}
}

//AddProductToCart
func AddProductToCart(userid string, productid string, client *mongo.Client) (string, error) {
	collection := client.Database("cartdb").Collection("cart")
	curlist, err := GetCartItemByUserId(userid, client)
	if err != nil {
		log.Fatalf("error geting cartitem :%v", err)
	}
	newlist := curlist.ProductId
	newlist = append(newlist, productid)
	_, err = collection.UpdateOne(context.TODO(), bson.D{{"_id", userid}}, bson.D{{"$set", bson.D{{"productid", newlist}}}})
	if err != nil {
		log.Fatalf("error adding item to cart: %v", err)
		return string(""), err
	} else {
		return string("sucessfullly updated"), nil
	}

}

//DeleteCartByUserId
func DeleteCartByUserId(userid string, client *mongo.Client) (string, error) {
	collection := client.Database("cartdb").Collection("cart")

	_, err := collection.DeleteOne(context.TODO(), bson.D{{"_id", userid}})
	if err != nil {
		log.Fatalf("err getting cart : %v", err)
		return string(""), err
	}
	return string("sucessfully deleted"), nil

}

// DeleteCartItemByProductId
func DeleteCartItemByProductId(userid, productid string, client *mongo.Client) (string, error) {
	collection := client.Database("cartdb").Collection("cart")
	cart, err := GetCartItemByUserId(userid, client)
	if err != nil {
		log.Fatalf("error geting cartitem :%v", err)
	}
	curlist := cart.ProductId
	newlist := []string{}
	for i := 0; i < len(curlist); i++ {
		if curlist[i] != productid {
			newlist = append(newlist, curlist[i])
		}
	}
	_, err = collection.UpdateOne(context.TODO(), bson.D{{"_id", userid}}, bson.D{{"$set", bson.D{{"productid", newlist}}}})
	if err != nil {
		log.Fatalf("error adding item to cart: %v", err)
		return string(""), err
	} else {
		return string("sucessfullly updated"), nil
	}
}
