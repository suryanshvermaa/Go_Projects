package controllers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suryanshvermaa/restaurant-management/database"
	"github.com/suryanshvermaa/restaurant-management/helpers"
	"github.com/suryanshvermaa/restaurant-management/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func GetFoods() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func GetFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := ctx.Param("food_id")
		var food models.Food
		defer cancel()
		err := foodCollection.FindOne(c, bson.M{"food_id": foodId}).Decode(&food)
		if err != nil {
			helpers.NewError(ctx, 400, err.Error())
			return
		}
		helpers.JsonResponse(ctx, 200, "food fetched successfully", food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func UpdateFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func round(num float64) int {

}

func toFixed(num float64, precision int) float64 {

}
