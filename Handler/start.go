package create

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Startserver() {
	server := Start()
	port := 0
	if err := Serv(server, port); err != nil {
		log.Fatalf("Server encountered an error: %v", err)
	}
}

func Start() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Homehandler)
	return mux
}

func Serv(handler http.Handler, port int) error {
	envPort := os.Getenv("PORT")
	if envPort != "" {
		fmt.Sscanf(envPort, "%d", &port)
	}
}
		