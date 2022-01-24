package api

import (
	"fmt"
	"net/http"
	"student-management/model"
	
	"github.com/gin-gonic/gin"
)

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
	if err := context.ShouldBind(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("req is %v\n", req)
	result := make([]model.Student, 0)
	values, err := server.store.FilterStudents(req.Gender, req.Age, req.Nrc)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	for _, v := range values {
		result = append(result, v)
	}
	
	context.JSON(http.StatusOK, result)
}

func (server *Server) GetDetailStudent(context *gin.Context) {
	var req model.GetDetailStudentRequest
	if err := context.ShouldBindUri(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "enter student id"})
		return
	}
	student, err := server.store.GetDetailStudent(int64(req.Id))
	if err != nil {
		return
	}
	
	context.JSON(http.StatusOK, student)
}
