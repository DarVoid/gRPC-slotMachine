package game

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/darvoid/gRPC-slotMachine/slotMachine"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) CreateGame(ctx context.Context, newGame *CreateGameRequest) (*NewGameReply, error) {
	log.Printf("Received message from client: WinChance:%v, TotalJogadas:%v\n", newGame.GetWinChance(), newGame.GetTotalJogadas())
	slot, err := slotMachine.Setup(int(newGame.GetTotalJogadas()), int(newGame.GetWinChance()))
	if err != nil {
		return nil, errors.New("error creating game")
	}
	fmt.Printf("%v\n", slot.CheckGameState())
	return &NewGameReply{GameId: slot.ID}, nil
}

func (s *Server) PlayGame(ctx context.Context, play *PlayRequest) (*ResultPlayReply, error) {
	log.Printf("Received message from client: GameId: %v, Player Name:%v, Lucky Quote:%v\n", play.GetGameId(), play.GetName(), play.GetLuckyQuote())
	now := time.Now()
	_, found := slotMachine.Games[play.GetGameId()]
	if !found {
		return nil, errors.New("game not found")
	}
	rward, err := slotMachine.Games[play.GetGameId()].Play(
		slotMachine.Person{Name: play.GetName(), LuckyQuote: play.GetLuckyQuote(), LastPlayed: now})
	log.Printf("Game State: %v\n", slotMachine.Games[play.GetGameId()].CheckGameState())

	if err != nil {
		log.Printf("Error trying to play Game: %v\n", err)
	}

	return &ResultPlayReply{GameId: play.GetGameId(), Name: play.GetName(), LuckyQuote: play.GetLuckyQuote(), Reward: rward}, nil
}
func (s *Server) GameExists(ctx context.Context, req *ShowGameRequest) (*GameExistsReply, error) {
	log.Printf("Received message from client: GameId: %v\n", req.GetGameId())
	_, found := slotMachine.Games[req.GetGameId()]

	if found {
		return &GameExistsReply{Exists: true}, nil
	}

	return &GameExistsReply{Exists: false}, nil
}
