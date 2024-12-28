package create

import (
	"database/sql"
	"log"
	"os"
)

func ToEquipdb() *sql.DB {
	if _, err := os.Stat("./database"); os.IsNotExist(err) {
		err := os.Mkdir("./database", 0o755)
		if err != nil {
			log.Fatal("filde to create dirictory ", err)
		}
	}
	db, err := sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		log.Fatal("fole to open sql database", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("filde to crete sql database", err)
	}
	return db
}
