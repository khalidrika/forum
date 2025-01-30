package create

import (
	_ "github.com/mattn/go-sqlite3"
)

func Run() {
	if !Initialise() {
		return
	}
	Startserver()
}
