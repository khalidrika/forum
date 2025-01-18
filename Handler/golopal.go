package create

import (
	"database/sql"
)

var (
	DB    *sql.DB
	liink *link
)

type link struct {
	Errorpage string `json:"error"`	
}
