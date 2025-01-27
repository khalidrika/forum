package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// path for login in ueser
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Errorhandler(w, http.StatusMethodNotAllowed, "Method Not Allowd", "Only Post Allowd Method ")
		return
	}
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("error decodinig respos body ", err)
		Errorhandler(w, http.StatusBadRequest, "Bad Request", "Invalid request body")
		return
	}
	if user.Email == "" || user.Username == "" || user.Password == "" {
		log.Println("masing requerd failds")
		Errorhandler(w, http.StatusBadRequest, "Bad Request", "Invalid server ere request")
		return
	}
	fmt.Println(user)
	// hash password
	hashengPasswor, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("failde to hashing password", err)
		Errorhandler(w, http.StatusInternalServerError, "Internel server error", "filde to hash password")
		return
	}
	// memorese user from database
	query := `INSERT INTO user (email, username, paswword, created_at) VALUES (?, ?, ?, ?)`
	_, err = DB.Exec(query, user.Email, user.Username, string(hashengPasswor), time.Now())
	if err != nil {
		log.Println("error isiting user to database", err)
		Errorhandler(w, http.StatusConflict, "Conflicet", "Email or Username aleady exists")
		return
	}
	log.Println("tam tasjil", user.Email)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("user registered successfully"))
}

// final

func ServeJS(w http.ResponseWriter, r *http.Request) {
	fileJS := "." + r.URL.Path
	filejs, err := os.ReadFile(fileJS)
	if err != nil {
		Errorhandler(w, http.StatusForbidden, http.StatusText(http.StatusForbidden), "error")
		return
	}
	http.ServeContent(w, r, fileJS, time.Now(), bytes.NewReader(filejs))
}

func ServeFile(w http.ResponseWriter, r *http.Request) {
	style := "." + r.URL.Path

	file, err := os.Open(style)
	if err != nil {
		http.Error(w, "faiulde to open file style", 404)
		return
	}
	http.ServeContent(w, r, file.Name(), time.Now(), file)
}

func Serv(handler http.Handler, port int) error {
	envPort := os.Getenv("PORT")
	if envPort != "" {
		fmt.Sscanf(envPort, "%d", &port)
	}
	addr := fmt.Sprintf(":%d", port)
	listner, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
		return err
	}
	defer listner.Close()

	server := &http.Server{
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		ErrorLog:     log.Default(),
	}
	log.Printf("starting server on http://localhost:%d", listner.Addr().(*net.TCPAddr).Port)
	return server.Serve(listner)
}
