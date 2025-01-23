package create

import (
	"bytes"
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
	mux.HandleFunc("/style/", ServeFile)
	mux.HandleFunc("/js/", ServeJS)
	return mux
}

func ServeJS(w http.ResponseWriter, r *http.Request) {
	fileJS := "." + r.URL.Path
	filejs, err := os.ReadFile(fileJS)
	if err != nil {
		Errorhandler(w, http.StatusForbidden, http.StatusText(http.StatusForbidden), "error")
		return
	}
	http.ServeContent(w, r, fileJS, time.Now(), bytes.NewReader(filejs))
}

func ServeFile(w http.ResponseWriter, r *http.Request) {
	style := "." + r.URL.Path

	file, err := os.Open(style)
	if err != nil {
		http.Error(w, "faiulde to open file style", 404)
		return
	}
	http.ServeContent(w, r, file.Name(), time.Now(), file)
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
