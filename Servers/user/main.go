package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	postgresclientmodel "github.com/capslock-inc/gprc-demo/Protos/database/postgresclient"
	usermodel "github.com/capslock-inc/gprc-demo/Protos/users"
	core "github.com/capslock-inc/gprc-demo/Servers/user/Core"
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
	p := flag.Int("port", 8405, "cartserver port")
	flag.Parse()
	port := fmt.Sprintf(":%d", *p)
	rpcserver, con := grpcclientforpostgres()
	defer con.Close()
	conn, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening to port : %v", err)
	}
	Userrpc := grpc.NewServer()
	usermodel.RegisterUserServiceServer(Userrpc, &core.UserServer{
		PostgresRPC: rpcserver,
	})

	log.Printf("server listening to %v", conn.Addr())
	if err := Userrpc.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
