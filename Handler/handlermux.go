package create

import (
	"bytes"
	"html/template"
	"net/http"
	"strings"
)

func Homehandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Errorhandler(w, 404, "page not fande", "looks like youre lost!")
		return
	}
	if r.Method != http.MethodGet {
		Errorhandler(w, 405, "method not allowd", "Only GET method is allowd!")
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
		Errorhandler(w, 500, "Internel Server Error", "Somthing seems wrong, try again later")
		return
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		if strings.HasPrefix(filename, "error.html") {
			ServClodeEroor(w, data.(EroorData), err)
			return
		}
		Errorhandler(w, 500, "Internal Server Error", "Somthing seems worng, try again leter")
		return
	}
	buf.WriteTo(w)
}
