package infra

import (
	"database/sql"
	"log"

	// driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

func Connect() *sql.DB {
	db, err := sql.Open("pgx", "user=user password=password host=postgres-test port=5432 dbname=postgres sslmode=disable timezone=Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}
