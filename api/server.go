package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"student-management/middleware"
	"student-management/store"
	"student-management/utils"
)

type Server struct {
	store  store.Store
	router *gin.Engine
}

func NewServer(store store.Store) Server {
	router := gin.Default()
	server := Server{store: store, router: router}
	authRouter := router.Group("/api")
	authRouter.Use(middleware.AuthMiddleware())
	authRouter.POST("/register", server.CreateStudent)
	authRouter.GET("/students", server.GetStudents)
	authRouter.GET("/student/:id", server.GetDetailStudent)
	authRouter.GET("/some", func(c *gin.Context) {
		studentId, ok := c.Get(utils.STUDENT_ID)
		if ok {
			c.JSON(http.StatusOK, gin.H{"msg": studentId})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "authentication error"})
		}
	})

	return server
}

func (server Server) Start() error {
	return server.router.Run()
}
