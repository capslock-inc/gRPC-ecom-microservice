package core

import (
	"log"

	model "github.com/capslock-inc/gprc-demo/Servers/postgresclient/Model"
	"gorm.io/gorm"
)

func InsertProduct(data *model.RpcProduct, db *gorm.DB) error {
	err := db.Create(&model.Product{
		Product_Name:     data.Name,
		Product_Price:    data.Price,
		Product_descript: data.Descript,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
func GetProduct(id int, db *gorm.DB) (*model.Product, error) {
	var retriveddata model.Product
	err := db.First(&retriveddata, id).Error
	if err != nil {
		return nil, err
	}
	return &retriveddata, nil
}
func GetAllProduct(db *gorm.DB) ([]model.Product, error) {
	var products []model.Product
	err := db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
func UpdateProduct(id int, data *model.RpcProduct, db *gorm.DB) error {
	result, err := GetProduct(id, db)
	if err != nil {
		return err
	}
	result.Product_Name = data.Name
	result.Product_Price = data.Price
	result.Product_descript = data.Descript

	if err = db.Save(result).Error; err != nil {
		log.Fatalf("error updating data: %v", err)
	}
	return nil
}
func DeleteProduct(id int, db *gorm.DB) error {
	err := db.Delete(&model.Product{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func InsertUser(data *model.RpcUser, db *gorm.DB) error {
	if err := db.Create(&model.User{
		Name:     data.Name,
		Password: data.Password,
	}).Error; err != nil {
		log.Fatalf("error inserting User : %v", err)
		return err
	}
	return nil
}
func GetUser(id int, db *gorm.DB) (*model.User, error) {
	var data model.User
	err := db.First(&data, id).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
func GetAllUser(db *gorm.DB) ([]model.User, error) {
	var users []model.User
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil

}
func UpdateUser(id int, data *model.RpcUser, db *gorm.DB) error {
	result, err := GetUser(id, db)
	if err != nil {
		return err
	}
	result.Name = data.Name
	result.Password = data.Password
	if err := db.Save(&result).Error; err != nil {
		log.Fatalf("error upadating user")
		return err
	}
	return nil
}
func DeleteUser(id int, db *gorm.DB) error {
	err := db.Delete(&model.User{}, id).Error
	if err != nil {
		return err
	}
	return err
}
