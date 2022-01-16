package api

import (
	"github.com/gin-gonic/gin"
	"student-management/store"
)

type Server struct {
	store  store.Store
	router *gin.Engine
}

func NewServer(store store.Store) Server {
	router := gin.Default()
	server := Server{store: store, router: router}
	router.POST("/register", server.CreateStudent)
	router.GET("/students", server.GetStudents)
	router.GET("/student/:id", server.GetDetailStudent)

	return server
}

func (server Server) Start() error {
	return server.router.Run()
}
