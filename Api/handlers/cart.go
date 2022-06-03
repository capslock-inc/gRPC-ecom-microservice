package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/capslock-inc/gprc-demo/Api/model"
	cartmodel "github.com/capslock-inc/gprc-demo/Protos/cart"
	"google.golang.org/grpc"
)

type Cart struct {
	l *log.Logger
}

func NewCart(logs *log.Logger) *Cart {
	return &Cart{
		l: logs,
	}
}

func (s *Cart) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.l.Printf("Cart visited 🥳 by %s with method: %s", r.RemoteAddr, r.Method)
	if r.Method == "GET" {
		s.get(w, r)
	}

}
func (s *Cart) cartclientrpc() cartmodel.CartServiceClient {
	addr := ":8403"
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		s.l.Fatalf("error Dialing to cartgrpcserver 😥 : %v", err)
	}
	client := cartmodel.NewCartServiceClient(conn)
	return client
}

func (s *Cart) get(w http.ResponseWriter, r *http.Request) {
	rpcclient := s.cartclientrpc()

	req := &model.UserId{}
	err := req.FromJSON(r.Body)
	if err != nil {
		s.l.Fatalf("error json marshal ")
	}

	result, err := rpcclient.GetCartItem(context.TODO(), &cartmodel.Userid{
		Id: strconv.Itoa(req.UID),
	})
	id, _ := strconv.Atoi(result.GetId())
	res := &model.CartItem{
		UID:   id,
		Plist: result.GetProductidlist(),
	}
	err = res.ToJson(w)
	if err != nil {
		w.Write([]byte("error"))
	}
}
