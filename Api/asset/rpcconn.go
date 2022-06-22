package asset

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	cartmodel "github.com/capslock-inc/gprc-demo/Protos/cart"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func CartClientRPC(l *log.Logger) (cartmodel.CartServiceClient, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	addr := fmt.Sprintf("%s:8403", os.Getenv("CART"))
	l.Println(addr)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		l.Printf("error Dialing to cartgrpcserver 😥 : %v", err)
		return nil, err
	}
	client := cartmodel.NewCartServiceClient(conn)
	return client, nil
}
