package create

import (
	"fmt"
	"os"
)

func Handlerconection(D *ports) {
	port.Getport()
	fmt.Println("Port:", ports.Port)
	fmt.Println("API Port:", ports.api)
}

func (p *ports) Getport() {
	p.port = os.Getenv("PORT")
	p.api = os.Getenv("APIPORT")
}

// func Homehandler(w http.ResponseWriter, r *http.Request) {
// 	return
// }
