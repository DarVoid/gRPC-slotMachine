package slotMachine

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	guuid "github.com/google/uuid"
)

//instances of players
type Person struct {
	Name       string
	LastPlayed time.Time
	LuckyQuote string
	reward     bool
	Resume     string
}

//pool of rewards
type Game struct {
	entries   []Person
	jogadas   int
	perc      int
	totalWins int
	last      int
	ID        string
}

//returns some of the fields as string (private)
func (p Person) toString() string {
	return fmt.Sprintf("Name: %v\nLast Played: %v\nLuckyQuote: %v\n", p.Name, p.LastPlayed, p.LuckyQuote)
}

//starts a new game with no players
func Setup(numberPlayers int, chanceWinning int) (*Game, error) {
	if chanceWinning > 100 {
		return nil, fmt.Errorf("Learn Math please")
	}
	numero := 100
	if numberPlayers >= 100 {
		numero = numberPlayers
	}
	game := &Game{entries: make([]Person, numero)}
	game.jogadas = numero
	game.perc = chanceWinning
	game.totalWins = 0
	game.ID = guuid.New().String()
	return game, nil

}

//executes a play by a player on a game
func (g *Game) Play(p Person) (bool, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	if n <= int64(g.perc) {
		p.reward = true
		g.totalWins++
	}

	if g.totalWins > g.jogadas*g.perc/100 {
		p.reward = false
		g.totalWins--

	}
	g.last++
	if g.last > len(g.entries) {
		g.last--
		return false, errors.New("Numero de Jogadas excedidas")
	} else {
		g.entries[g.last-1] = p
	}
	return p.reward, nil
}

//returns a few info about the game
func (g Game) CheckGameState() string {
	return fmt.Sprintf("\nGame ID: %v\nTotal de Jogadas: %v\nWinning chance: %v\nTotal Wins: %v\nLast Element: %v\n", g.ID, g.jogadas, g.perc, g.totalWins, g.last)
}

//returns numbeer of current wins
func (g Game) GetTotalVictories() int {
	return g.totalWins
}
