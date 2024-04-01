package hello

import (
	context "context"
	"log"
)

type Server struct{}

func (s *Server) Talkie(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Received message from client: %s", in.Content)
	return &Message{Content: "from golang server"}, nil
}

func (s *Server) mustEmbedUnimplementedHelloServiceServer() {

}
