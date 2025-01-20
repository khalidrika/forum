package create

import (
	"fmt"
	"net/http"
)

func Startserver() {
	server := Start()
	fmt.Println(server)
}

func Start() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Homehandler)
	return mux
}
