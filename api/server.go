package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	router.POST("token", func(context *gin.Context) {
		token, err := utils.CreateToken(1)
		if err != nil {
			//log.Fatal("create token failed")    // follow by a call to os.Exit(1) , so this state will terminate without response to client
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "failed to create token"})
		}
		context.JSON(http.StatusOK, gin.H{"token": token})
	})
	authRouter := router.Group("/api")
	//authRouter.Use(middleware.AuthMiddleware())
	authRouter.POST("/register", server.CreateStudent)
	authRouter.GET("/students", server.GetStudents)
	authRouter.GET("/student/:id", server.GetDetailStudent)

	return server
}

func (server Server) Start() error {
	return server.router.Run()
}
