package create

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Run() {
	if !Initialise() {
		return
	}
	// port.Getport()
	// Inilink()
}

func Initialise() bool {
	if InitialisePort() {
		return false
	}
	InitialiseLink()
	InitialiseDB()
	return true
}

func InitialisePort() bool {
	port := flag.Bool("print-port", false, "print a random available port and exit")
	flag.Parse()
	if *port {
		listner, err := net.Listen("tcp", ":0")
		if err != nil {
			fmt.Fprintf(os.Stderr, "failde to find a random port: %v\n", err)
			os.Exit(1)
		}
		defer listner.Close()
		fmt.Println(listner.Addr().(*net.TCPAddr).Port)
		return true
	}
	return false
}

func InitialiseLink() {
	cont, err := os.ReadFile("./json/clod.json")
	if err != nil {
		log.Printf("failde to red clod.json:%v\n", err)
		return
	}
	err = json.Unmarshal(cont, &liink)
	if err != nil {
		log.Printf("failde to parse clod.json: %v\n", err)
		return
	}
	if liink.Errorpage == "" {
		log.Printf("warning errorpage list is empty")
	}
}

func InitialiseDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		log.Fatal("failde open to sqlt3 database: \n", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("failed", err)
	}
	conten, err := os.ReadFile("./database/sqlt.sql")
	if err != nil {
		log.Fatal("failde to read dfile sqlite", err)
		fmt.Println(conten)
	}

	if _, err := DB.Exec(string(conten)); err != nil {
		log.Fatal("Failed to create database tables:", err)
	}
}
