package create

import (
	"database/sql"
)

var (
	port  ports
	DB    *sql.DB
	liink link
)

type link struct {
	errorpage string `json:"error`
}

type ports struct {
	port string
	api  string
}

type clainte struct {
	username string
	Email    string
	pasword  string
}
