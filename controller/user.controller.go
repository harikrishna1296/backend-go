package controller

import (
	"fmt"
	"net/http"
	"server/database"
	"server/models"
	"strings"

	"github.com/gin-gonic/gin"
)

type GetUserInfoStruct struct {
	Email string
}

func GetUserInfo(context *gin.Context) {
	var body GetUserInfoStruct
	type User models.User
	var user User
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	database.Instance.Find(&user, User{Email: body.Email})
	fmt.Println("record", user, body)
	context.JSON(http.StatusOK, gin.H{"record": "Success", "data": user})
}

func CreateUser(context *gin.Context) {
	var body models.User
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := body.HashPassword(body.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	usernameReplace := strings.ReplaceAll(body.Name, " ", "_")
	body.Username = strings.ToLower(usernameReplace)
	fmt.Println(body)
	record := database.Instance.Create(&body)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"user": body.Name, "email": body.Email})
}
