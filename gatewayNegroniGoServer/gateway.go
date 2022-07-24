package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/darvoid/gRPC-slotMachine/game"
	"github.com/darvoid/gRPC-slotMachine/gameVitaeDashboard"

	"google.golang.org/grpc"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func main() {

	//create gorilla router instance
	router := mux.NewRouter()

	//create routes
	router.HandleFunc("/create", CreateGameHandle).Methods("POST")
	router.HandleFunc("/play", PlayGameHandle).Methods("POST")
	router.HandleFunc("/exists", GameExistsHandle).Methods("POST")
	router.HandleFunc("/retrieve-session-data", HandleRetrieval).Methods("POST")
	router.HandleFunc("/", SetCORHeader).Methods("OPTIONS") //default route in case of no match

	n := negroni.Classic()                     // new negroni instance with default middleware
	wrapped := n.With()                        // add additional middleware funcs here
	wrapped.UseHandler(router)                 // link gorilla router instance to negroni instance
	recovery := negroni.NewRecovery()          // setup recovery strategy
	recovery.PanicHandlerFunc = reportToSentry // setup function to be called in case of panic
	wrapped.Use(recovery)                      // link negroni instance with recory strategy
	handler := cors.New(cors.Options{
		AllowedOrigins:     []string{"http://localhost:4200"},
		AllowedMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE"},
		AllowedHeaders:     []string{"Accept, Accept-Language,Authorization, Content-Type, YourOwnHeader"},
		AllowCredentials:   true,
		OptionsPassthrough: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	}).Handler(wrapped) // create http.handler with negroni instance
	http.ListenAndServe(":8080", handler) // serve
}

//report to Sentry

func reportToSentry(info *negroni.PanicInformation) {

	fmt.Println("bad things happened")

}

func SetCORHeader(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language,Authorization, Content-Type, YourOwnHeader")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}

	// https://github.com/gorilla/mux/blob/master/example_cors_method_middleware_test.go
	if r.Method == http.MethodOptions {
		return
	}

}

type userSessionRequest struct {
	User      string
	PageIndex int64
	PageSize  int64
	OrderBy   string
	Asc       bool
}

// http method handlers

func HandleRetrieval(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language,Authorization, Content-Type, YourOwnHeader")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}
	var conn *grpc.ClientConn
	//vars := mux.Vars(r) //path parameters
	//r.ParseForm() //query parameters
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect%v\n", err)
	}
	defer conn.Close()

	c := gameVitaeDashboard.NewGameVitaeServiceClient(conn)
	var req userSessionRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Request: %v\n", req)
	response, err := c.RetrieveSessionData(context.Background(), &gameVitaeDashboard.SessionParameterRequest{
		User: req.User,

		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		OrderBy:   req.OrderBy,
		Asc:       req.Asc,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err2 := fmt.Fprintf(w, fmt.Sprintf(err.Error(), http.StatusBadRequest))
		if err2 != nil {
			log.Printf("error writing to response")
		}
		return
	}
	log.Printf("Successfully retrieved data for user: %v\n", req.User)
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
	w.Header().Set("Content-Type", "application/json")
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
	log.Printf("Game with ID: %v played\n", gameToPlay.GetGameId())
	val, err := json.Marshal(response)
	if err != nil {
		log.Printf("error marshalling")
	}
	w.Header().Set("Content-Type", "application/json")
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
	log.Printf("Game with ID: %v verified\n", gameExists.GetGameId())
	val, err := json.Marshal(response)
	if err != nil {
		log.Printf("error marshalling")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = fmt.Fprintf(w, string(val))
	if err != nil {
		log.Printf("error writing to response")
	}
}
