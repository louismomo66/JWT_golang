package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/louismomo66/JWT_golang/controllers"
	"github.com/louismomo66/JWT_golang/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users",controller.GetUsers())
	incomingRoutes.GET("/user/:user_id",controller.GetUser())
}