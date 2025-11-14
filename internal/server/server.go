package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ehharvey/chat-app/internal/abc"
	"github.com/ehharvey/chat-app/internal/config"
	"github.com/spf13/viper"
)

type Server struct {
	Mux        *http.ServeMux
	abcService *abc.Service
}

func newServer(abcService *abc.Service) *Server {
	result := Server{
		Mux:        http.NewServeMux(),
		abcService: abcService,
	}

	result.createRoutes()

	return &result
}

func (s *Server) createRoutes() {
	s.Mux.HandleFunc("/", s.HandleIndex)
	s.Mux.HandleFunc("/foo", s.HandleCreateFoo)
}

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func (s *Server) HandleCreateFoo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	var input abc.InsertOneFooParams
	inputParseErr := json.NewDecoder(r.Body).Decode(&input)

	if inputParseErr != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	result := s.abcService.CreateFoo(r.Context(), input)

	if len(result.ValidationResultAggregate) > 0 {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if result.InternalServerError != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func StartServer() {
	// Config initialization
	config.InitializeConfig()
	config.LoadConfig()

	// Telemetry: TODO

	// Database:

	// Services

	// Server:
	server := Server{}
	server.createRoutes()

	port := viper.GetString("server.port")
	log.Fatal(http.ListenAndServe(":"+port, server.Mux))
}
