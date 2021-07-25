package repository

import (
	"database/sql"
	"fmt"
	"os"

	log "github.com/golang/glog"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	log.Info("Connecting to postgres instance at:", psqlInfo)

	dbtmp, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	r.db = dbtmp

	err = r.db.Ping()
	if err != nil {
		panic(err)
	}

	log.Info("Successfully connected!")
	return nil
}
