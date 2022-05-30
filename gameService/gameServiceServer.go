package main

import (
	"fmt"
	"log"
	"net"

	"github.com/darvoid/gRPC-slotMachine/game"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on pport 9000: %v\n", err)
	}
	fmt.Println("WE LISTENING BOYYYYY!!")

	s := game.Server{}
	grpcServer := grpc.NewServer()
	game.RegisterGameServiceServer(grpcServer, &s)
	fmt.Println("WE GAMING BOYYYYY!!")

	//shows available resources on gRPC server
	reflection.Register(grpcServer)
	fmt.Println("Here's what we got")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve:  %v\n", err)
	}
}
