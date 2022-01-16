package api

import (
	"net/http"
	"student-management/model"

	"github.com/gin-gonic/gin"
)

//request handlers for student
func (server *Server) CreateStudent(context *gin.Context) {

	// check if nrc is already registered
	var student model.Student
	if err := context.ShouldBindJSON(&student); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := server.store.CreateStudent(student)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "create student success"})

}

func (server *Server) GetStudents(context *gin.Context) {
	var req model.GetFilteredStudentRequest
	if err := context.BindQuery(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, server.store.GetStudents(req.Gender))
}

func (server *Server) GetDetailStudent(context *gin.Context) {
	var req model.GetDetailStudentRequest
	if err := context.ShouldBindUri(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "enter student id"})
		return
	}

	context.JSON(http.StatusOK, server.store.GetDetailStudent(req.Id))
}
