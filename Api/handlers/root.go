package handlers

import (
	"log"
	"net/http"
)

type Root struct {
	l *log.Logger
}

func NewRoot(l *log.Logger) *Root {
	return &Root{l}
}

func (s *Root) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.l.Printf("🥳 Root visited by %s with method: %s", r.RemoteAddr, r.Method)
	w.Write([]byte(`/Cart 👉 post(Create cart){"uid":int},
put(add to cart){"uid":int,"pid":int},
delete(delete item from cart){"uid":int,"pid":int}.

/Cart/uid 👉 get(get item from cart){},
delete(delete cart){}. 
	`))
}
