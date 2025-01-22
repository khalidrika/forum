package create

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
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
	addr := fmt.Sprintf(":%d", port)
	listner, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
		return err
	}
	defer listner.Close()

	server := &http.Server{
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		ErrorLog:     log.Default(),
	}
	log.Printf("starting server on http://localhost:%d", listner.Addr().(*net.TCPAddr).Port)
	return server.Serve(listner)
}
