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
