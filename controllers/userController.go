package controllers

import (
"context"
"log"
"fmt"
"strconv"
"time"
"net/http"
"github.com/gin-gonic/gin"
"github.com/go-playground/validator/v10"
helper "github.com/louismomo66/JWT_golang/helpers"
"github.com/louismomo66/JWT_golang/models"
"github.com/louismomo66/JWT_golang/helpers"
"github.com/louismomo66/JWT_golang/database"
"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, *user)
var validate = validator.New()

func HashPassword()

func VeryfyPassword()

func SignUp()

func GetUsers()

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id")

	if err :=	helper.MatchUserTypeToUid(c, userId); err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
	}
}