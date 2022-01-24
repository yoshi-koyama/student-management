package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"student-management/model"
	"student-management/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := model.AuthHeader{}
		
		if err := c.ShouldBindHeader(&header); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "You don't have permission!"})
			return
		}
		
		claim, err := utils.ValidateToken(header.Token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "You don't have permission!"})
			return
		}
		
		studentId := claim[utils.STUDENT_ID]
		c.Set(utils.STUDENT_ID, studentId)
		
		c.Next()
	}
}
