package model

import (
	"encoding/json"
	"io"
)

type Productid struct {
	UID int `json:"uid"`
	PID int `json:"pid"`
}

func (p *Productid) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)

}
func (p *Productid) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

type CartItem struct {
	UID   int      `json:"uid"`
	Plist []string `json:"Productlist"`
}

func (p *CartItem) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)

}
func (p *CartItem) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

type Userid struct {
	UID int `json:"uid"`
}

func (p *Userid) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)

}
func (p *Userid) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
