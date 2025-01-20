package create

import (
	"html/template"
	"net/http"
)

func Homehandler(w http.ResponseWriter, r http.Request) {
	if r.URL.Path != "/" {
		Erorhandler(w, 404, "page not fande")
		return
	}
	if r.Method != http.MethodGet {
		Erorhandler(w, 405, "method not alowd")
		return
	}
	Parsehome(w, "", "./template/index.html")
}

func Parsehome(w http.ResponseWriter, data any, filename string) {
	tmpl, err := template.ParseFiles(filename)
    if err != nil {
        
        
    }
}
