package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	cartmodel "github.com/capslock-inc/gprc-demo/Protos/cart"
	mongoclientmodel "github.com/capslock-inc/gprc-demo/Protos/database/mongoclient"
	core "github.com/capslock-inc/gprc-demo/Servers/cart/Core"
	"google.golang.org/grpc"
)

func grpcclientformongoclient() (mongoclientmodel.MongoClientServiceClient, *grpc.ClientConn) {
	rpcserveraddress := "localhost:8401"

	conn, err := grpc.Dial(
		rpcserveraddress,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("mongoclientserver is down 😥: %v", err)
	}
	rpcclient := mongoclientmodel.NewMongoClientServiceClient(conn)

	return rpcclient, conn

}

func main() {
	p := flag.Int("port", 8403, "cartserver port")
	flag.Parse()
	port := fmt.Sprintf(":%d", *p)
	rpcserver, con := grpcclientformongoclient()
	defer con.Close()
	conn, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening to port : %v", err)
	}
	Cartrpc := grpc.NewServer()
	cartmodel.RegisterCartServiceServer(Cartrpc, &core.CartServer{
		Mongorpc: rpcserver,
	})

	log.Printf("server listening to %v", conn.Addr())
	if err := Cartrpc.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
