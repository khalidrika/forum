package create

import (
	"html/template"
	"net/http"
	"strings"
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
	ParseAndEx(w, "", "./template/index.html")
}

func ParseAndEx(w http.ResponseWriter, data any, filename string) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		if strings.HasPrefix(filename, "error.html") {
			ServClodeEroor(w, data.(EroorData), err)
			return
		}
	}
}
