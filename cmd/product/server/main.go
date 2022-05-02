package main

import (
	"log"
	"net"

	productcore "github.com/capslock-inc/gprc-demo/ProductService/product-core"
	product_model "github.com/capslock-inc/gprc-demo/ProductService/product-model"
	"google.golang.org/grpc"
)

func main() {
	port := ":5000"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	s := grpc.NewServer()

	product_model.RegisterProductServiceServer(s, &productcore.ProductServer{})
	log.Printf("server listing at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
