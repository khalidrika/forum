package create

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func Run() {
	if !Initialise() {
		return
	}
	// port.Getport()
	// Inilink()
}

func Initialise() bool {
	if Initialiseport() {
		return false
	}
	return true
}

func Initialiseport() bool {
	portt := flag.Bool("print-port", false, "print a random avilable portand exit")
	flag.Parse()
	if *portt {
		lidtner, err := net.Listen("tcp", ":0")
		if err != nil {
			fmt.Fprintf(os.Stderr, "failde to find a random port: %v\n", err)
			os.Exit(1)
		}
		defer lidtner.Close()
		fmt.Println(lidtner.Addr().(*net.TCPAddr).Port)
		return true
	}
	return false
}

// db := ip.ToEquipdb()
// defer db.Close()
// http.HandleFunc("/", ip.Homehandler)
// fmt.Println("page http://localhost:8087")
// log.Fatal(http.ListenAndServe(":8087", nil))
