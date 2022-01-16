package main

import (
	"log"
	"student-management/api"
	"student-management/store"
)

func main() {
	store := store.NewStore()
	server := api.NewServer(store)

	if err := server.Start(); err != nil {
		log.Fatal("cannot start")
	}
}
