package main

import (
	"log"

	"github.com/darvoid/gRPC-slotMachine/game"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect%v\n", err)
	}
	defer conn.Close()

	c := game.NewGameServiceClient(conn)

	message := game.CreateGameRequest{
		WinChance: 20, TotalJogadas: 200,
	}

	response, err := c.CreateGame(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when creating game %v", err)
	}
	log.Printf("Game with ID: %v created\n", response.GameId)
}
