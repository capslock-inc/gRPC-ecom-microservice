package asset

import (
	"log"

	cartmodel "github.com/capslock-inc/gprc-demo/Protos/cart"
	productmodel "github.com/capslock-inc/gprc-demo/Protos/product"
	usermodel "github.com/capslock-inc/gprc-demo/Protos/users"
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
func ProductClientRPC(l *log.Logger) productmodel.ProductServiceClient {
	addr := ":8404"
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		l.Fatalf("error Dialing to cartgrpcserver 😥 : %v", err)
	}
	client := productmodel.NewProductServiceClient(conn)
	return client
}
func UserClientRPC(l *log.Logger) usermodel.UserServiceClient {
	addr := ":8405"
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		l.Fatalf("error Dialing to cartgrpcserver 😥 : %v", err)
	}
	client := usermodel.NewUserServiceClient(conn)
	return client
}
