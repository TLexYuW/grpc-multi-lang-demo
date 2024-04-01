package service

import (
	"context"
	"go_grpc_demo/pb"
	"log"
)

type MessageSenderServerImpl struct {
	*pb.UnimplementedMessageSenderServer
}

func (MessageSenderServerImpl) Send(context context.Context, request *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Println("receive message:", request.GetSaySomething())
	resp := &pb.MessageResponse{}
	resp.ResponseSomething = "Response From gRPC Server"
	return resp, nil
}
