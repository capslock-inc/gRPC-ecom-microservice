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
	w.Write([]byte(`/Cart 👉 get,post,put,delete 
/Porduct 👉  get,post,put,delete
/User 👉  get,post,put,delete
	`))
}
