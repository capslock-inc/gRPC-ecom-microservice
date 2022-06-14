package main

import (
	"context"
	"flag"
	"net"
	"time"

	"fmt"
	"log"

	mongoclientmodel "github.com/capslock-inc/gprc-demo/Protos/database/mongoclient"
	core "github.com/capslock-inc/gprc-demo/Servers/mongoclient/Core"
	"github.com/capslock-inc/gprc-demo/logmodel"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
)

func DbINIT(port int) (*mongo.Client, context.Context) {

	//loading env in current executing dir
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	// parseing mongodb url
	dburl := fmt.Sprintf("mongodb://root:password@localhost:%d", port)

	// creating new client for mongodb
	client, err := mongo.NewClient(options.Client().ApplyURI(dburl))
	if err != nil {
		log.Fatalf("newclientreturnerror:%v", err)
	}

	// setting  timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// establishing db connection
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("clientconnecterror:%v", err)
	}

	// verifying established connection by pinging
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("connection establised")
	}
	return client, ctx

}

func main() {
	// -port flag in terminal
	mongoport := flag.Int("mongoport", 27017, "mongodb port")
	// parsing flag data
	flag.Parse()

	// logger
	logs := logmodel.Logger("MongoClient Server 👉 ")
	// initiating mongodb
	client, ctx := DbINIT(*mongoport)

	// grpc server init
	serverport := flag.Int("port", 8401, "server port")
	port := fmt.Sprintf(":%d", *serverport)
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error initiating net.listen : %v", err)
	}
	grpcserver := grpc.NewServer()
	mongoclientmodel.RegisterMongoClientServiceServer(grpcserver, &core.MongoClientServer{Client: client, Logs: logs})
	logs.Printf("🚀 server listening to %v", listen.Addr())
	if err := grpcserver.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// terminating db connection when program is done
	defer client.Disconnect(ctx)
}
