package main

import (
	"context"
	"demo/grpc/protobuf/user"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8889", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := user.NewUserServiceClient(conn)

	req := &user.UserRequest{
		Id: 1,
	}

	response, _ := client.GetUser(context.Background(), req)

	resp, err := json.Marshal(response)

	fmt.Printf("%s", resp)

}
