package main

import (
	"context"
	"flag"
	"time"

	"fmt"
	"log"
	"os"

	core "github.com/capslock-inc/gprc-demo/Servers/mongoclient/Core"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DbINIT(port int) (*mongo.Client, context.Context) {

	//loading env in current executing dir
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	// parseing mongodb url
	dburl := fmt.Sprintf("mongodb://%s:%s@localhost:%d", os.Getenv("ADMIN"), os.Getenv("PASSWORD"), port)

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
	p := flag.Int("port", 27017, "mongodb port")
	// parsing flag data
	flag.Parse()
	port := *p

	// initiating mongodb
	client, ctx := DbINIT(port)

	// create new cart for user
	// statement, err := core.CreateCart(&model.CartItem{
	// 	UserId: "test2",
	// 	ProductId: []string{
	// 		"firstitem",
	// 		"seconditem",
	// 	},
	// }, client)
	// if err != nil {
	// 	log.Fatalf("error writing : %v", err)
	// } else {
	// 	fmt.Printf("%s", statement)
	// }

	arr := core.GetCart(client)
	fmt.Println(arr)

	// terminating db connection when program is done
	defer client.Disconnect(ctx)
}
