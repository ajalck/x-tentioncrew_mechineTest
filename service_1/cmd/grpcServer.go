package main

import (
	"log"
	"net"

	"github.com/ajalck/service_1/pkg/config"
	service "github.com/ajalck/service_1/pkg/grpcService"
	"github.com/ajalck/service_1/pkg/pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func StartGrpcServer(config *config.Config, db *gorm.DB) {
	lis, err := net.Listen("tcp", config.GrpcPort)
	if err != nil {
		log.Fatalf("Failed to load env: %v", err)
	}
	grpcServer := grpc.NewServer()

	server := &service.UserServer{
		DB: db,
	}
	pb.RegisterUsersServer(grpcServer, server)

	log.Println("gRPC Server started listening at: ", config.GrpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve gRPC server: ", err)
	}
}
