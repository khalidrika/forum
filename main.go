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
	http.HandleFunc("/", ip.Homehandler)
	fmt.Println("page http://localhost:8087")
	log.Fatal(http.ListenAndServe(":8087", nil))
}
