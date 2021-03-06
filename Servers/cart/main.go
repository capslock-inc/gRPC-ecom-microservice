package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	cartmodel "github.com/capslock-inc/gprc-demo/Protos/cart"
	mongoclientmodel "github.com/capslock-inc/gprc-demo/Protos/database/mongoclient"
	core "github.com/capslock-inc/gprc-demo/Servers/cart/Core"
	"github.com/capslock-inc/gprc-demo/logmodel"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func grpcclientformongoclient() (mongoclientmodel.MongoClientServiceClient, *grpc.ClientConn) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	rpcserveraddress := fmt.Sprintf("%s:8401", os.Getenv("MONGOCLIENT"))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	log.Println(rpcserveraddress)
	conn, err := grpc.DialContext(
		ctx,
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

	logs := logmodel.Logger("Cart Server 👉 ")

	rpcserver, con := grpcclientformongoclient()
	defer con.Close()

	conn, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening to port : %v", err)
	}

	Cartrpc := grpc.NewServer()
	cartmodel.RegisterCartServiceServer(Cartrpc, &core.CartServer{
		Mongorpc: rpcserver,
		Logs:     logs,
	})

	logs.Printf("🚀 server listening to %v", conn.Addr())
	if err := Cartrpc.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
