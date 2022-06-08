package core

import (
	"context"
	"fmt"
	"log"
	"strconv"

	postgresclientmodel "github.com/capslock-inc/gprc-demo/Protos/database/postgresclient"
	model "github.com/capslock-inc/gprc-demo/Servers/postgresclient/Model"
	"gorm.io/gorm"
)

// 	rpc AllUser(emptyparam) returns (userlist);
//  rpc UserById(userid) returns (user);
//     rpc AddUser(user) returns (status);
//     rpc UpdateUser (user) returns (status);
//     rpc DeleteUser(userid) returns (status);

//     rpc AllProduct(emptyparam) returns (productlist);
//     rpc ProductById(productid) returns (product);
//     rpc AddProduct(product) returns (status);
//     rpc UpdateProduct (product) returns (status);
//     rpc DeleteProduct(productid) returns (status);

type PostgresClientServer struct {
	postgresclientmodel.UnimplementedPostgresClientServiceServer
	Db *gorm.DB
}

func (s *PostgresClientServer) AllUser(ctx context.Context, _ *postgresclientmodel.Emptyparam) (*postgresclientmodel.Userlist, error) {
	var users *postgresclientmodel.Userlist
	result, err := GetAllUser(s.Db)
	if err != nil {
		log.Printf("errorgeting all user :%v ", err)
		return nil, err
	}
	for _, user := range result {
		users.List = append(users.List,
			&postgresclientmodel.User{
				Id:       strconv.Itoa(user.Id),
				Name:     user.Name,
				Password: user.Password,
			},
		)
	}
	return users, nil
}

func (s *PostgresClientServer) UserById(ctx context.Context, id *postgresclientmodel.Userid) (*postgresclientmodel.User, error) {
	a, err := strconv.Atoi(id.Id)
	if err != nil {
		log.Printf("error converting string to int")
	}
	result, err := GetUser(a, s.Db)
	if err != nil {
		return nil, err
	}
	return &postgresclientmodel.User{
		Id:       strconv.Itoa(result.Id),
		Name:     result.Name,
		Password: result.Password,
	}, nil
}
func (s *PostgresClientServer) AddUser(ctx context.Context, data *postgresclientmodel.User) (*postgresclientmodel.Status, error) {
	err := InsertUser(&model.RpcUser{
		Name:     data.Name,
		Password: data.Password,
	}, s.Db)
	if err != nil {
		return &postgresclientmodel.Status{
			Value: fmt.Sprintf("error: %v", err),
		}, err

	}
	return &postgresclientmodel.Status{
		Value: "sucessfully created",
	}, nil
}
func (s *PostgresClientServer) UpdateUser(ctx context.Context, data *postgresclientmodel.User) (*postgresclientmodel.Status, error) {
	id, err := strconv.Atoi(data.Id)
	if err != nil {
		return nil, err
	}
	err = UpdateUser(id, &model.RpcUser{
		Name:     data.Name,
		Password: data.Password,
	}, s.Db)
	if err != nil {
		return nil, err
	}
	return &postgresclientmodel.Status{
		Value: "sucessfully updated",
	}, nil
}
func (s *PostgresClientServer) DeleteUser(ctx context.Context, data *postgresclientmodel.Userid) (*postgresclientmodel.Status, error) {
	id, err := strconv.Atoi(data.Id)
	if err != nil {
		return nil, err
	}
	err = DeleteUser(id, s.Db)
	if err != nil {
		return nil, err
	}
	return &postgresclientmodel.Status{
		Value: "sucessfully deleted",
	}, nil
}

// producthandlers
func (s *PostgresClientServer) AllProduct(ctx context.Context, _ *postgresclientmodel.Emptyparam) (*postgresclientmodel.Productlist, error) {
	var productlist postgresclientmodel.Productlist
	result, err := GetAllProduct(s.Db)
	if err != nil {
		return nil, err
	}
	for _, product := range result {
		productlist.List = append(productlist.List, &postgresclientmodel.Product{
			Id:          strconv.Itoa(product.Id),
			Name:        product.Product_Name,
			Price:       strconv.Itoa(product.Product_Price),
			Description: product.Product_descript,
		})
	}
	return &productlist, nil
}
func (s *PostgresClientServer) ProductById(ctx context.Context, data *postgresclientmodel.Productid) (*postgresclientmodel.Product, error) {
	id, err := strconv.Atoi(data.Id)
	if err != nil {
		return nil, err
	}
	result, err := GetProduct(id, s.Db)
	if err != nil {
		return nil, err
	}
	return &postgresclientmodel.Product{
		Id:          strconv.Itoa(result.Id),
		Name:        result.Product_Name,
		Price:       strconv.Itoa(result.Product_Price),
		Description: result.Product_descript,
	}, nil
}

func (s *PostgresClientServer) AddProduct(ctx context.Context, data *postgresclientmodel.Product) (*postgresclientmodel.Status, error) {
	price, err := strconv.Atoi(data.Price)
	if err != nil {
		return nil, err
	}
	err = InsertProduct(&model.RpcProduct{
		Name:     data.Name,
		Price:    price,
		Descript: data.Description,
	}, s.Db)
	if err != nil {
		return nil, err
	}
	return &postgresclientmodel.Status{
		Value: "sucessfully inserted",
	}, nil
}

func (S *PostgresClientServer) UpdateProduct(ctx context.Context, data *postgresclientmodel.Product) (*postgresclientmodel.Status, error) {

	id, err := strconv.Atoi(data.Id)
	if err != nil {
		return nil, err
	}
	price, err := strconv.Atoi(data.Price)
	if err != nil {
		return nil, err
	}
	err = UpdateProduct(id, &model.RpcProduct{
		Name:     data.Name,
		Price:    price,
		Descript: data.Description,
	}, S.Db)
	if err != nil {
		return nil, err
	}
	return &postgresclientmodel.Status{
		Value: "sucessfully update",
	}, nil

}
func (s *PostgresClientServer) DeleteProduct(ctx context.Context, id *postgresclientmodel.Productid) (*postgresclientmodel.Status, error) {

	pid, err := strconv.Atoi(id.Id)
	if err != nil {
		return nil, err
	}
	err = DeleteProduct(pid, s.Db)
	if err != nil {
		return nil, err
	}
	return &postgresclientmodel.Status{
		Value: "sucessfully deleted",
	}, nil

}
