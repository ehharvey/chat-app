package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ehharvey/chat-app/internal/config"
	"github.com/spf13/viper"
)

func StartServer() {
	// Config initialization
	config.InitializeConfig()
	config.LoadConfig()

	// Telemetry: TODO

	// Database: TODO

	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloServer)

	port := viper.GetString("server.port")
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
