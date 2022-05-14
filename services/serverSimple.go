package main

import (
	"log"
	"net"

	"github.com/darvoid/slot/game"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on pport 9000: %v\n", err)
	}

	s := game.Server{}
	grpcServer := grpc.NewServer()
	game.RegisterGameServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve:  %v\n", err)
	}
}
