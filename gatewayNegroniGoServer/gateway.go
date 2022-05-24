package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/darvoid/gRPC-slotMachine/game"

	"google.golang.org/grpc"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func main() {
	//mux

	router := mux.NewRouter()

	router.HandleFunc("/create", CreateGameHandle).Methods("POST")
	router.HandleFunc("/play", PlayGameHandle).Methods("POST")
	router.HandleFunc("/exists", GameExistsHandle).Methods("POST")
	//router

	n := negroni.Classic() // Includes some default middlewares
	wrapped := n.With()    // insert new middleware funcs here

	wrapped.UseHandler(router)
	recovery := negroni.NewRecovery()
	recovery.PanicHandlerFunc = reportToSentry
	wrapped.Use(recovery)
	//n
	handler := cors.Default().Handler(wrapped)
	http.ListenAndServe(":8080", handler)
}

//report to Sentry

func reportToSentry(info *negroni.PanicInformation) {

	fmt.Println("bad things happened")

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
	if err != nil {
		log.Printf("error marshalling")
	}
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, string(val))
	if err != nil {
		log.Printf("error writing to response")
	}

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
		_, err = fmt.Fprintf(w, fmt.Sprintf(err.Error(), http.StatusBadRequest))
		if err != nil {
			log.Printf("error writing to response")
		}
	}
	log.Printf("Game with ID: %v played\n", response.GameId)
	val, err := json.Marshal(response)
	if err != nil {
		log.Printf("error marshalling")
	}
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, string(val))
	if err != nil {
		log.Printf("error writing to response")
	}
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
		fmt.Fprintf(w, "could not connect to service\n", err)
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
	if err != nil {
		log.Printf("error marshalling")
	}
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, string(val))
	if err != nil {
		log.Printf("error writing to response")
	}
}
