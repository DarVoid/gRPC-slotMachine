//this package is a demo for handling in-memory rigged slot machine logic
package slotMachine

//package which handles game logic of rigged slot machine

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
	Name       string    //name of the player for each play
	LastPlayed time.Time //date of each play
	LuckyQuote string    //lucky quote for each play
	reward     bool      //saves reward true/false
	Resume     string    //base64 resume string of all the other data
}

//Game with of rewards
type Game struct {
	entries   []Person
	jogadas   int
	perc      int
	totalWins int
	last      int
	ID        string // Id of game
}

//returns some of the fields as string
func (p Person) ToString() string {
	return fmt.Sprintf("Name: %v\nLast Played: %v\nLuckyQuote: %v\n", p.Name, p.LastPlayed, p.LuckyQuote)
}

//in memory array of games
var Games = make(map[string]*Game)

//starts a new game with no players
func Setup(numberPlayers int, chanceWinning int) (*Game, error) {
	if chanceWinning > 100 {
		return nil, fmt.Errorf("learn math please")
	}
	if chanceWinning == 0 {
		return nil, fmt.Errorf("no win game")
	}
	if numberPlayers == 0 {
		return nil, fmt.Errorf("unplayable")
	}
	numero := 100
	if numberPlayers >= 100 { //this can be commented if we don't wanna enforce min plays
		numero = numberPlayers
	}
	game := Game{entries: make([]Person, numero)}
	game.ID = guuid.New().String()
	game.jogadas = numero
	game.perc = chanceWinning

	Games[game.ID] = &game
	return &game, nil

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
		return false, errors.New("numero de jogadas excedidas")
	} else {
		g.entries[g.last-1] = p
	}
	g.CheckGameState()
	return p.reward, nil
}

//returns a few info about the game
func (g Game) CheckGameState() string {
	return fmt.Sprintf("\nGame ID: %v\nTotal de Jogadas: %v\nWinning chance: %v\nTotal Wins: %v\nLast Element: %v\n", g.ID, g.jogadas, g.perc, g.totalWins, g.last)
}

//returns a few info about the game
func (g Game) OutputCheckGameState() string {
	return fmt.Sprintf("Total de Jogadas: %v\nWinning chance: %v\nTotal Wins: %v\nLast Element: %v\n", g.jogadas, g.perc, g.totalWins, g.last)
}

//returns numbeer of current wins
func (g Game) GetTotalVictories() int {
	return g.totalWins
}

//returns an array of all games
func ListGamesInMemory() ([]Game, error) {

	gameArray := []Game{}

	for key, value := range Games {
		fmt.Println("Key:", key, "Value:", value)
		gameArray = append(gameArray, *value)
	}
	return gameArray, nil
}

//returns an in-memory game with given id
func ShowGame(id string) (Game, error) {

	game := Games[id]
	return *game, nil
}
