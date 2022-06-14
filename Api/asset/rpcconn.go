package asset

import (
	"context"
	"log"
	"time"

	cartmodel "github.com/capslock-inc/gprc-demo/Protos/cart"
	"google.golang.org/grpc"
)

func CartClientRPC(l *log.Logger) (cartmodel.CartServiceClient, error) {
	addr := ":8403"

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
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
