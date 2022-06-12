package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/capslock-inc/gprc-demo/Api/model"
	cartmodel "github.com/capslock-inc/gprc-demo/Protos/cart"
	"github.com/capslock-inc/gprc-demo/logmodel"
)

type CartWithParams struct {
	l       *log.Logger
	cartrpc cartmodel.CartServiceClient
}

func NewCartWithParams(logs *log.Logger, rpccon cartmodel.CartServiceClient) *CartWithParams {
	return &CartWithParams{
		l:       logs,
		cartrpc: rpccon,
	}
}

func (s *CartWithParams) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	s.l.Printf("CartWithParams visited 🥳 by %s with method: %s", r.RemoteAddr, r.Method)

	switch r.Method {
	case "GET":
		s.get(w, r)
	case "DELETE":
		s.delete(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("cant find"))
	}

}

func (s *CartWithParams) get(w http.ResponseWriter, r *http.Request) {

	rpcclient := s.cartrpc
	UID := strings.TrimPrefix(r.URL.Path, "/cart/")

	jdata, err := json.Marshal(UID)
	if err != nil {
		s.l.Println(logmodel.Logdata("101", "😥", err.Error(), "error marshalling json in GET: /cart/"))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return
	}

	result, err := rpcclient.GetCartItem(context.TODO(), &cartmodel.Userid{
		Id: UID,
	})
	if err != nil {
		s.l.Println(logmodel.Logdata("101", "😥", err.Error(), "error retriving data in GET: /cart/"))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return
	}

	id, _ := strconv.Atoi(result.GetId())
	res := &model.CartItem{
		UID:   id,
		Plist: result.GetProductidlist(),
	}

	err = res.ToJson(w)
	if err != nil {
		s.l.Println(logmodel.Logdata("101", "😥", err.Error(), "error TOJSON in GET: /cart/"))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return
	}

	s.l.Println(logmodel.Logdata("101", "😃", string(jdata), "GET: /cart/ successfull"))
}

func (s *CartWithParams) delete(w http.ResponseWriter, r *http.Request) {

	rpcclient := s.cartrpc
	UID := strings.TrimPrefix(r.URL.Path, "/cart/")

	jdata, err := json.Marshal(UID)
	if err != nil {
		s.l.Println(logmodel.Logdata("102", "😥", err.Error(), "error marshalling json in DELETE: /cart/"))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return
	}

	_, err = rpcclient.DeleteCart(context.TODO(), &cartmodel.Userid{
		Id: UID,
	})
	if err != nil {
		s.l.Println(logmodel.Logdata("102", "😥", err.Error(), "error deleting in DELETE: /cart/"))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id: /cart/id"))
		w.Write([]byte("Error deleting"))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Successfully deleted"))
	s.l.Println(logmodel.Logdata("102", "😃", string(jdata), "DELETE: /cart/ successfull"))
}

type Cart struct {
	l       *log.Logger
	cartrpc cartmodel.CartServiceClient
}

func NewCart(logs *log.Logger, rpccon cartmodel.CartServiceClient) *Cart {
	return &Cart{
		l:       logs,
		cartrpc: rpccon,
	}
}

func (s *Cart) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	s.l.Printf("Cart visited 🥳 by %s with method: %s", r.RemoteAddr, r.Method)

	switch r.Method {
	case "POST":
		s.post(w, r)
	case "PUT":
		s.put(w, r)
	case "DELETE":
		s.delete(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("cant find"))
	}
}
func (s *Cart) post(w http.ResponseWriter, r *http.Request) {

	userid := &model.Userid{}
	err := userid.FromJSON(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid Body:: body should be {'uid': int}"))
		s.l.Printf(logmodel.Logdata("103", "😥", err.Error(), "error FromJson in POST: /cart"))
	}

	jdata, err := json.Marshal(userid)
	if err != nil {
		s.l.Println(logmodel.Logdata("103", "😥", err.Error(), "error marshalling json in POST: /cart"))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return
	}

	s.cartrpc.CreateNewCart(context.TODO(), &cartmodel.Userid{
		Id: strconv.Itoa(userid.UID),
	})

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Successfully created"))
	s.l.Println(logmodel.Logdata("103", "😃", string(jdata), "POST: /cart successfull"))
}

func (s *Cart) put(w http.ResponseWriter, r *http.Request) {

	productid := &model.Productid{}
	err := productid.FromJSON(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid Body:: body should be {'uid': int,'pid':int }"))
		s.l.Printf(logmodel.Logdata("104", "😥", err.Error(), "error FromJson in PUT: /cart"))
	}

	jdata, err := json.Marshal(productid)
	if err != nil {
		s.l.Println(logmodel.Logdata("103", "😥", err.Error(), "error marshalling json in POST: /cart"))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return
	}

	s.cartrpc.AddToCart(context.TODO(), &cartmodel.Productid{
		Id:     strconv.Itoa(productid.PID),
		Userid: strconv.Itoa(productid.UID),
	})

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Successfully added to cart"))
	s.l.Println(logmodel.Logdata("103", "😃", string(jdata), "POST: /cart successfull"))

}
func (s *Cart) delete(w http.ResponseWriter, r *http.Request) {

	productid := &model.Productid{}
	err := productid.FromJSON(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid Body:: body should be {'uid': int,'pid':int }"))
		s.l.Printf("error retriving data")
	}

	s.l.Printf("retrived data : %#v", productid)
	s.cartrpc.DeleteCartItem(context.TODO(), &cartmodel.Productid{
		Id:     strconv.Itoa(productid.PID),
		Userid: strconv.Itoa(productid.UID),
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("product deleted from cart successfully"))
}
