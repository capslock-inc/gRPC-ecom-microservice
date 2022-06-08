package model

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
