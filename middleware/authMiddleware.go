package middleware

import (
	"fmt"
	"net/http"
	helper "github.com/louismomo66/JWT_golang/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc{
	return func(c *gin.Context){
		clientToken := c.Request.Header.Get("token")
		if clientToken == ""{
			c.JSON(http.StatusInternalServerError,gin.H{"error":fmt.Sprintf("No Authorisation header provided")})
			c.Abort()
			return
		}

		claims,err := helper.ValidateToken(clientToken)

	}
}