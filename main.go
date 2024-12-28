package main

import (
	"fmt"
	"log"
	"net/http"

	ip "forum/create"
)

func main() {
	db := ip.ToEquipdb()
	defer db.Close()
	http.HandleFunc("home", ip.Homehandler)
	log.Panicln("page http://localhost:8087")
	if err := http.ListenAndServe(":8087", nil); err != nil {
		fmt.Println(err.Error())
	}
}
