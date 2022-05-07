package model

type CartItem struct {
	UserId    string   `bson:"_id"`
	ProductId []string `bson:"productid"`
}
