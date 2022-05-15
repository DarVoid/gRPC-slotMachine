package game

import (
	"fmt"
	"log"
	"time"

	"github.com/darvoid/slot/slotMachine"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) CreateGame(ctx context.Context, newGame *CreateGameRequest) (*NewGameReply, error) {
	log.Printf("Received message from client: WinChance:%v, TotalJogadas:%v\n", newGame.GetWinChance(), newGame.GetTotalJogadas())
	slot, err := slotMachine.Setup(int(newGame.GetTotalJogadas()), int(newGame.GetWinChance()))
	if err != nil {
		log.Fatalf("Error creating Game: %v\n", err)
	}
	fmt.Printf("%v\n", slot.CheckGameState())
	return &NewGameReply{GameId: slot.ID}, nil
}

func (s *Server) PlayGame(ctx context.Context, play *PlayRequest) (*ResultPlayRequest, error) {
	log.Printf("Received message from client: GameId: %v, Player Name:%v, Lucky Quote:%v\n", play.GetGameId(), play.GetName(), play.GetLuckyQuote())
	now := time.Now()
	rward, err := slotMachine.Games[play.GetGameId()].Play(
		slotMachine.Person{Name: play.GetName(), LuckyQuote: play.GetLuckyQuote(), LastPlayed: now})
	log.Printf("Game State: %v\n", slotMachine.Games[play.GetGameId()].CheckGameState())

	if err != nil {
		log.Printf("Error trying to play Game: %v\n", err)
	}
	fmt.Printf("%v\n", play)

	return &ResultPlayRequest{GameId: play.GetGameId(), Name: play.GetName(), LuckyQuote: play.GetLuckyQuote(), Reward: rward}, nil
}
func (s *Server) GameExists(ctx context.Context, req *ShowGameRequest) (*GameExistsReply, error) {
	log.Printf("Received message from client: GameId: %v\n", req.GetGameId())
	err := slotMachine.Games[req.GetGameId()]

	if err != nil {
		return &GameExistsReply{Exists: false}, nil
	}

	return &GameExistsReply{Exists: true}, nil
}
