package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	routes "github.com/louismomo66/JWT_golang/routes"
)

func main(){
	err := godotenv.Load(".env")
if err != nil{
	log.Fatal("Error loading .env file")
}
port := os.Getenv("PORT")

if port == ""{
	port= "9000"
}

router := gin.New()
router.Use(gin.Logger())

routes.AuthRoutes(router)
routes.UserRoutes(router)
 
router.GET("/api-1", func(c *gin.Context){
	c.JSON(200,gin.H{"success":"Access granted for api-1"})
})
router.GET("/api-2", func(c *gin.Context){
	c.JSON(200,gin.H{"success":"Access granted for api-2"})
})

router.Run(":" + port)
}