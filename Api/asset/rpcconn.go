package asset

import (
	"log"

	cartmodel "github.com/capslock-inc/gprc-demo/Protos/cart"
	"google.golang.org/grpc"
)

func CartClientRPC(l *log.Logger) cartmodel.CartServiceClient {
	addr := ":8403"
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		l.Fatalf("error Dialing to cartgrpcserver 😥 : %v", err)
	}
	client := cartmodel.NewCartServiceClient(conn)
	return client
}
