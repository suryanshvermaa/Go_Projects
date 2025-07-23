package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/suryanshvermaa/Go_Projects/go-mongo-jwt/database"
	"github.com/suryanshvermaa/Go_Projects/go-mongo-jwt/helpers"
	"github.com/suryanshvermaa/Go_Projects/go-mongo-jwt/utils"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(password *string)
func VerifyPassword()
func Signup()
func Login()

func GetUsers()
func GetUser() *gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helpers.MatchUserTypeToUid(c, userId); err != nil {
			utils.Response(c, http.StatusBadRequest, err.Error(), nil)
			return
		}

	}
}
