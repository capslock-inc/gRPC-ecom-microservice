package main

import (
	"log"

	mongoclientmodel "github.com/capslock-inc/gprc-demo/Protos/database/mongoclient"
	"google.golang.org/grpc"
)

func grpcclientformongoclient() *mongoclientmodel.MongoClientServiceClient {
	rpcserveraddress := "localhost:8401"

	conn, err := grpc.Dial(
		rpcserveraddress,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	defer conn.Close()
	if err != nil {
		log.Fatalf("unable to dial to rpcserver : %v", err)
	}
	rpcclient := mongoclientmodel.NewMongoClientServiceClient(conn)

	return &rpcclient

}

func main() {
}
