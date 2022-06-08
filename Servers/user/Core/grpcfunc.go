package core

import (
	"context"
	"log"

	postgresclientmodel "github.com/capslock-inc/gprc-demo/Protos/database/postgresclient"
	usermodel "github.com/capslock-inc/gprc-demo/Protos/users"
)

type UserServer struct {
	usermodel.UnimplementedUserServiceServer
	PostgresRPC postgresclientmodel.PostgresClientServiceClient
}

// rpc GetAllUser(emptyparam) returns (userlist);
// rpc GetUserById(userid) returns (user);
// rpc AddNewUser(user) returns (status);
// rpc UpdateUser (user) returns (status);
// rpc DeleteUser(userid) returns (status)

func (s *UserServer) GetAllUser(ctx context.Context, _ *usermodel.Emptyparam) (*usermodel.Userlist, error) {
	result, err := s.PostgresRPC.AllUser(ctx, &postgresclientmodel.Emptyparam{})
	if err != nil {
		log.Println("error getting all user")
		return nil, err
	}
	var list []*usermodel.User
	for _, l := range result.List {
		list = append(list, &usermodel.User{
			Id:       l.Id,
			Name:     l.Name,
			Password: l.Password,
		})
	}
	log.Println("getalluser sucessfull 😇")
	return &usermodel.Userlist{
		List: list,
	}, nil
}

func (s *UserServer) GetUserById(ctx context.Context, id *usermodel.Userid) (*usermodel.User, error) {
	result, err := s.PostgresRPC.UserById(
		ctx,
		&postgresclientmodel.Userid{
			Id: id.Id,
		},
	)
	if err != nil {
		log.Println("error getuserbyid")
		return nil, err
	}
	log.Println("getuserbyid successful 😇")
	return &usermodel.User{
		Id:   result.Id,
		Name: result.Name,
	}, nil

}

func (s *UserServer) AddNewUser(ctx context.Context, data *usermodel.User) (*usermodel.Status, error) {
	_, err := s.PostgresRPC.AddUser(
		ctx,
		&postgresclientmodel.User{
			Id:       data.Id,
			Name:     data.Name,
			Password: data.Password,
		},
	)
	if err != nil {
		log.Println("error addser")
		return nil, err
	}
	log.Println("adduser sucessful 😇")
	return &usermodel.Status{
		Value: "✔️ added successfully 😃",
	}, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, data *usermodel.User) (*usermodel.Status, error) {
	_, err := s.PostgresRPC.UpdateUser(
		ctx,
		&postgresclientmodel.User{
			Id:       data.Id,
			Name:     data.Name,
			Password: data.Password,
		},
	)
	if err != nil {
		log.Println("error upadteuser")
		return nil, err
	}
	log.Println("updateuser sucessful 😇")
	return &usermodel.Status{
		Value: "updated successfully 😃",
	}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, data *usermodel.Userid) (*usermodel.Status, error) {
	_, err := s.PostgresRPC.DeleteUser(
		ctx,
		&postgresclientmodel.Userid{
			Id: data.Id,
		},
	)
	if err != nil {
		log.Println("error deleteuser")
		return nil, err
	}
	log.Println("deleteuser successful 😇")
	return &usermodel.Status{
		Value: "deleted successfully 😃",
	}, nil
}
