package game

import (
	"fmt"
	"log"

	"github.com/darvoid/slot/slotMachine"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) CreateGame(ctx context.Context, newGame *CreateGameRequest) (*NewGameReply, error) {
	log.Printf("Received message from client: Winchance:%v, totalJogadas:%v\n", newGame.WinChance, newGame.TotalJogadas)
	slot, err := slotMachine.Setup(int(newGame.TotalJogadas), int(newGame.WinChance))
	if err != nil {
		log.Fatalf("Error creating Game: %v\n", err)
	}
	fmt.Printf("%v\n", slot)
	return &NewGameReply{GameId: 1}, nil
}
