package client

import (
	"log"

	"github.com/ajalck/service_1/pkg/pb"
	"google.golang.org/grpc"
)

type Service1Client struct {
	Client pb.UsersClient
}

func InitService1Client(url string) *Service1Client {
	cc, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Could not connect to product service client")
		grpc.WithReturnConnectionError()
	}
	return &Service1Client{
		Client: pb.NewUsersClient(cc),
	}
}


