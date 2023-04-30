package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hidekimva/golang/gateway/controllers"
)

func HandleRequest() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.GET("/status", controllers.Status)

	r.POST("/user", controllers.CreateUser)

	r.Run(":8081")

}
