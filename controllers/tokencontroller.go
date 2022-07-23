package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k-yoshigai/learning-jwt-go/auth"
	"github.com/k-yoshigai/learning-jwt-go/database"
	"github.com/k-yoshigai/learning-jwt-go/models"
)

type TokenRequest struct {
	Email    string `json:"email`
	Password string `json:"password`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.Instance.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	credentilaError := user.CheckPassword(request.Password)
	if credentilaError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})

}
