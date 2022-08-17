package main

import (
	"net/http"

	"server/controller"
	"server/database"

	"github.com/gin-gonic/gin"
)

func main() {
	//Initial Router
	dcs := "user=postgres password=harihema dbname=service port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	database.Connect(dcs)
	database.Migrate()
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			// tokenString, err := helper.GenerateJWT("harikrish1296@gmail.com")
			// if err != nil {
			// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			// 	c.Abort()
			// 	return
			// }
			// c.JSON(http.StatusOK, gin.H{"message": "Success", "token": tokenString})
			c.JSON(http.StatusOK, gin.H{"message": "Success"})
		})
		api.POST("/login", controller.GetUserInfo)
		api.POST("/user", controller.CreateUser)
	}
	return router
}
