package create

import "net/http"

func Homehandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Erorhandler(w, r, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	
}
