package create

import (
	"log"
	"net/http"
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
	mux.HandleFunc("/style/", ServeFile)
	mux.HandleFunc("/js/", ServeJS)
	mux.HandleFunc("/register", RegisterHandler)
	mux.HandleFunc("/login", LoginHandler)
	return mux
}
