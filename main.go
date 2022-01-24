package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"student-management/api"
	"student-management/store"
)

func main() {
	
	//database, err := sql.Open("postgres", "postgres://docker:password@localhost:5433/school?sslmode=disable")
	database, err := sql.Open("postgres", "postgres://localhost:5432/school?sslmode=disable")
	if err != nil {
		log.Fatal("db open error", err)
	}
	
	//we should check connection success or not here , because sql.Open just create Sql object
	err = database.Ping()
	if err != nil {
		log.Fatal("db connection error", err)
	}
	
	store2 := store.NewSqlStore(database)
	server := api.NewServer(store2)
	
	if err := server.Start(); err != nil {
		log.Fatal("cannot start")
	}
	
}
