package create

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type EroorData struct {
	Msg1       string
	Msg2       string
	statusCode int
}

func Erorhandler(w http.ResponseWriter, code int, misage string) {
}

func ServClodeEroor(w http.ResponseWriter, error EroorData, err error) {
	log.Println(err)
	errBody, err := GetEroorPage()
	if err != nil {
		http.Error(w, http.StatusText(error.statusCode), error.statusCode)
		log.Println(err)
		return
	}
	errBody = strings.Replace(errBody, "{{.Msg1}}", error.Msg1)
}

func GetEroorPage() (string, error) {
	url := liink.Errorpage

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failde to fetch error page: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failde to read error page conten: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failde to read page cntent: %v", err)
	}
	return string(body), nil
}
