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
	authRouter := router.Group("/api")
	authRouter.POST("/students", server.CreateStudent)
	authRouter.GET("/students", server.GetStudents)
	authRouter.GET("/students/:id", server.GetDetailStudent)

	return server
}

func (server Server) Start() error {
	return server.router.Run()
}
