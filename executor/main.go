package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/livensmi1e/tiny-ide/executor/proto"
	"github.com/livensmi1e/tiny-ide/executor/server"
)

// TODO: using zerolog + read from envs + dockerize
func main() {
	lis, err := net.Listen("tcp", ":"+server.PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterExecutorServer(s, &server.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
