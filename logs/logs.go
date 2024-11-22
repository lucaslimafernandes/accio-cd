package logs

import (
	"database/sql"
	"log"

	"github.com/lucaslimafernandes/accio-cd/handler"
)

func InsertLog(db *sql.DB, l *handler.LogEntry) {

	query := `INSERT INTO logs (timestamp, level, status, message) VALUES (?, ?, ?, ?)`

	_, err := db.Exec(query, l.Timestamp, l.Level, l.Status, l.Message)
	if err != nil {
		log.Fatalf("erro ao inserir log: %v\n", err)
	}

}
