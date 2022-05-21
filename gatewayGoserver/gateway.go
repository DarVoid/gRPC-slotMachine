package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/darvoid/gRPC-slotMachine/game"
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

	var newGame game.CreateGameRequest

	err = json.NewDecoder(r.Body).Decode(&newGame)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("WinChance: %v\nTotalJogadas: %v\n", newGame.WinChance, newGame.TotalJogadas)

	response, err := c.CreateGame(context.Background(), &newGame)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err2 := fmt.Fprintf(w, fmt.Sprintf(err.Error(), http.StatusBadRequest))
		if err2 != nil {
			log.Printf("error writing to response")
		}
		return
	}
	log.Printf("Game with ID: %v created\n", response.GameId)
	val, err := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(val))

}
func PlayGameHandle(w http.ResponseWriter, r *http.Request) {
	var conn *grpc.ClientConn

	var gameToPlay game.PlayRequest

	err := json.NewDecoder(r.Body).Decode(&gameToPlay)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	conn, err = grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect%v\n", err)
	}
	defer conn.Close()

	c := game.NewGameServiceClient(conn)

	response, err := c.PlayGame(context.Background(), &gameToPlay)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err2 := fmt.Fprintf(w, fmt.Sprintf(err.Error(), http.StatusBadRequest))
		if err2 != nil {
			log.Printf("error writing to response")
		}
	}
	log.Printf("Game with ID: %v played\n", response.GameId)
	val, err := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(val))
}

func GameExistsHandle(w http.ResponseWriter, r *http.Request) {
	var conn *grpc.ClientConn
	var gameExists game.ShowGameRequest

	err := json.NewDecoder(r.Body).Decode(&gameExists)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	conn, err = grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect%v\n", err)
	}
	defer conn.Close()

	c := game.NewGameServiceClient(conn)

	response, err := c.GameExists(context.Background(), &gameExists)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err2 := fmt.Fprintf(w, fmt.Sprintf(err.Error(), http.StatusBadRequest))
		if err2 != nil {
			log.Printf("error writing to response")
		}
	}
	log.Printf("Game with ID: %v verified\n", response)
	val, err := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(val))
}
