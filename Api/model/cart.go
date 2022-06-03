package model

import (
	"encoding/json"
	"io"
)

type UserId struct {
	UID int `json:"uid"`
}
type Productid struct {
	UID int `json:"uid"`
	PID int `json:"pid"`
}
type CartItem struct {
	UID   int      `json:"uid"`
	Plist []string `json:"Productlist"`
}

func (p *UserId) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)

}
func (p *CartItem) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
