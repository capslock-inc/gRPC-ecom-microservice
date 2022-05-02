package main

import (
	"context"
	"log"
	"time"

	product_model "github.com/capslock-inc/gprc-demo/ProductService/product-model"
	"google.golang.org/grpc"
)

func main() {
	address := "localhost:5000"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connent: %v", err)
	}
	defer conn.Close()
	c := product_model.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_product = make(map[string]string)
	new_product["lap"] = "black,slim"
	new_product["keyboard"] = "rgb,metal"

	for name, desc := range new_product {
		r, err := c.AddProduct(ctx, &product_model.Product{Name: name, Description: desc})
		if err != nil {
			log.Fatalf("could not create user : %v", err)
		}
		log.Printf("id: %s", r.GetValue())
	}
}
