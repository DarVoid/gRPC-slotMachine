package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

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

	r.HandleFunc("/create/{totalJogadas}/{winChance}", CreateGameHandle).Methods("POST")
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
	vars := mux.Vars(r) //path parameters
	//r.ParseForm() //query parameters
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect%v\n", err)
	}
	defer conn.Close()

	c := game.NewGameServiceClient(conn)

	//winChance, err := strconv.ParseInt(r.FormValue("winChance"), 10, 32)
	winChance, err := strconv.ParseInt(vars["winChance"], 10, 32)
	fmt.Println(winChance, err, reflect.TypeOf(winChance))

	//totalJogadas, err := strconv.ParseInt(r.FormValue("totalJogadas"), 10, 32)
	totalJogadas, err := strconv.ParseInt(vars["totalJogadas"], 10, 32)
	fmt.Println(totalJogadas, err, reflect.TypeOf(totalJogadas))

	message := game.CreateGameRequest{
		WinChance: int32(winChance), TotalJogadas: int32(totalJogadas),
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
