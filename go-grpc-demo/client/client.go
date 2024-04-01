package main

import (
	"context"
	"go_grpc_demo/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMessageSenderClient(conn)
	resp, err := client.Send(context.Background(), &pb.MessageRequest{SaySomething: "Hello World From Client!"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("receive message:", resp.GetResponseSomething())
}
