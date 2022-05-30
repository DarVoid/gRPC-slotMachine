package main

import (
	"fmt"
	"log"
	"net"

	"github.com/darvoid/gRPC-slotMachine/gameVitaeDashboard"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Failed to listen on port 9001: %v\n", err)
	}
	fmt.Println("WE LISTENING BOYYYYY!!")

	s := gameVitaeDashboard.Server{}
	grpcServer := grpc.NewServer()
	gameVitaeDashboard.RegisterGameVitaeServiceServer(grpcServer, &s)

	//shows available resources on gRPC server
	reflection.Register(grpcServer)
	fmt.Println("GameVitaeBeingServed")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve:  %v\n", err)
	}
}
