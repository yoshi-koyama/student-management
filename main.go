package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"student-management/api"
	"student-management/store"
)

func main() {

	database, err := sql.Open("postgres", "postgres://docker:password@localhost:5433/school?sslmode=disable")
	if err != nil {
		log.Fatal("db open error", err)
		return
	}

	store := store.NewSqlStore(database)
	server := api.NewServer(store)

	if err := server.Start(); err != nil {
		log.Fatal("cannot start")
	}
}
