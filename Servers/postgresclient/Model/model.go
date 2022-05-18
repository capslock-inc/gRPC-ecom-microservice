package model

import (
	"gorm.io/gorm"
)

// rpc model
type RpcProduct struct {
	Id       string
	Name     string
	Price    int
	Descript string
}
type RpcUser struct {
	Id       string
	Name     string
	Password string
}

// gorm model
type User struct {
	gorm.Model
	Id       int
	Name     string
	Password string
}
type Product struct {
	gorm.Model
	Id               int
	Product_Name     string
	Product_Price    int
	Product_descript string
}
