package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	postgresclientmodel "github.com/capslock-inc/gprc-demo/Protos/database/postgresclient"
	productmodel "github.com/capslock-inc/gprc-demo/Protos/product"
	core "github.com/capslock-inc/gprc-demo/Servers/product/Core"
	"google.golang.org/grpc"
)

func grpcclientforpostgres() (postgresclientmodel.PostgresClientServiceClient, *grpc.ClientConn) {
	rpcserveraddress := "localhost:8402"

	conn, err := grpc.Dial(
		rpcserveraddress,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Printf("postgresclient is down 😥: %v", err)
	}
	rpcclient := postgresclientmodel.NewPostgresClientServiceClient(conn)

	return rpcclient, conn

}

func main() {
	p := flag.Int("port", 8404, "cartserver port")
	flag.Parse()
	port := fmt.Sprintf(":%d", *p)
	rpcserver, con := grpcclientforpostgres()
	defer con.Close()
	conn, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening to port : %v", err)
	}
	Productrpc := grpc.NewServer()
	productmodel.RegisterProductServiceServer(
		Productrpc,
		&core.ProductServer{
			PostgresRPC: rpcserver,
		},
	)

	log.Printf("🚀 server listening to %v", conn.Addr())
	if err := Productrpc.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
	}

}
