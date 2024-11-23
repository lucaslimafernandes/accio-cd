package logs

import (
	"database/sql"
	"log"
	"time"
)

type LogEntry struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	Status    string    `json:"status"`
	Message   string    `json:"message"`
}

func InsertLog(db *sql.DB, l *LogEntry) {

	query := `INSERT INTO logs (timestamp, level, status, message) VALUES (?, ?, ?, ?)`

	_, err := db.Exec(query, l.Timestamp, l.Level, l.Status, l.Message)
	if err != nil {
		log.Fatalf("erro ao inserir log: %v\n", err)
	}

}
