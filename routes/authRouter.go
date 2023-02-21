package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/louismomo66/JWT_golang/contollers"
)

func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("users/signup",controller.Signup())
	incomingRoutes.POST("users/login",controller.Login())
}