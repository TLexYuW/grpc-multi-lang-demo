package main

import (
	"go_grpc_demo/pb"
	"go_grpc_demo/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()
	pb.RegisterMessageSenderServer(srv, service.MessageSenderServerImpl{})
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = srv.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
