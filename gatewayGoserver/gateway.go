package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/darvoid/slot/game"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {

	//https://github.com/gorilla/mux
	r := mux.NewRouter()
	r.HandleFunc("/create", MiddlewareOptionsDefaultPolicy).Methods("OPTIONS") //.Methods(http.MethodOptions)
	r.HandleFunc("/play", MiddlewareOptionsDefaultPolicy).Methods("OPTIONS")   //.Methods(http.MethodOptions)
	r.HandleFunc("/exists", MiddlewareOptionsDefaultPolicy).Methods("OPTIONS") //.Methods(http.MethodOptions)

	r.HandleFunc("/create", CreateGameHandle).Methods("POST")
	r.HandleFunc("/play", PlayGameHandle).Methods("POST")
	r.HandleFunc("/exists", GameExistsHandle).Methods("POST")

	r.Use(mux.CORSMethodMiddleware(r))
	http.ListenAndServe(":8080", r)
}

func MiddlewareOptionsDefaultPolicy(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	// https://github.com/gorilla/mux/blob/master/example_cors_method_middleware_test.go
	if r.Method == http.MethodOptions {
		return
	}
}

type NewGame struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

type CreateGameRequest struct {
	WinChance    int32 `json:"winChance"`
	TotalJogadas int32 `json:"totalJogadas"`
}

func CreateGameHandle(w http.ResponseWriter, r *http.Request) {

	var conn *grpc.ClientConn
	//vars := mux.Vars(r) //path parameters
	//r.ParseForm() //query parameters
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect%v\n", err)
	}
	defer conn.Close()

	c := game.NewGameServiceClient(conn)

	var p CreateGameRequest

	errr := json.NewDecoder(r.Body).Decode(&p)
	if errr != nil {
		http.Error(w, errr.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("WinChance: %v\nTotalJogadas: %v\n", p.WinChance, p.TotalJogadas)

	message := game.CreateGameRequest{
		WinChance: p.WinChance, TotalJogadas: p.TotalJogadas,
	}

	response, err := c.CreateGame(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when creating game %v", err)
	}
	log.Printf("Game with ID: %v created\n", response.GameId)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, response.String())

}
func PlayGameHandle(w http.ResponseWriter, r *http.Request) {
	var conn *grpc.ClientConn
	vars := mux.Vars(r)

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect%v\n", err)
	}
	defer conn.Close()

	c := game.NewGameServiceClient(conn)

	message := game.PlayRequest{
		GameId: vars["gameId"], Name: vars["name"], LuckyQuote: vars["luckyQuote"],
	}

	response, err := c.PlayGame(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when playing game %v", err)
	}
	log.Printf("Game with ID: %v played\n", response.GameId)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, response.String())
}

func GameExistsHandle(w http.ResponseWriter, r *http.Request) {
	var conn *grpc.ClientConn
	vars := mux.Vars(r)

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect%v\n", err)
	}
	defer conn.Close()

	c := game.NewGameServiceClient(conn)

	message := game.ShowGameRequest{
		GameId: vars["gameId"],
	}

	response, err := c.GameExists(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when verifying game %v", err)
	}
	log.Printf("Game with ID: %v verified\n", response)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, response.String())
}
