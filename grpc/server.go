package main

import (
	"context"
	"demo/grpc/protobuf/user"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type UserServiceInterface interface {
	GetUser(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error)
}

type UserServiceStruct struct {
}

func NewUserService() UserServiceInterface {
	return &UserServiceStruct{}
}

func (userService *UserServiceStruct) GetUser(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error) {
	response := &user.UserResponse{
		Id:   req.Id,
		Name: "Hello World",
	}

	return response, nil
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8889")

	if err != nil {
		panic(err)
	}

	fmt.Println("listen on 127.0.0.1:8889")

	grpcServer := grpc.NewServer()

	var userService UserServiceInterface

	userService = NewUserService()
	user.RegisterUserServiceServer(grpcServer, userService)

	err = grpcServer.Serve(l)

	if err != nil {
		println(err)
	}

}
