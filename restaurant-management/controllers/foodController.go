package controllers

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/suryanshvermaa/restaurant-management/database"
	"github.com/suryanshvermaa/restaurant-management/helpers"
	"github.com/suryanshvermaa/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		recordPerPage, err := strconv.Atoi(ctx.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}
		var foods []bson.M
		page, err := strconv.Atoi(ctx.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}
		startIndex := (page - 1) * recordPerPage
		matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}}
		groupStage := bson.D{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: bson.D{{Key: "_id", Value: "null"}}},
			{Key: "total_count", Value: bson.D{{Key: "$sum", Value: 1}}},
			{Key: "data", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
		}}}
		projectStage := bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "total_count", Value: 1},
				{Key: "food_items", Value: bson.D{{Key: "$slice", Value: []interface{}{"$data", startIndex, recordPerPage}}}},
			}},
		}
		result, err := foodCollection.Aggregate(c, mongo.Pipeline{matchStage, groupStage, projectStage})
		if err != nil {
			helpers.NewError(ctx, 500, "error occured while listing food items")
			return
		}
		if err = result.All(c, &foods); err != nil {
			helpers.NewError(ctx, 500, err.Error())
			return
		}
		defer cancel()
		helpers.JsonResponse(ctx, 200, "foods fetched successfully", foods)
	}
}

func GetFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := ctx.Param("food_id")
		var food models.Food
		err := foodCollection.FindOne(c, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()
		if err != nil {
			helpers.NewError(ctx, 400, err.Error())
			return
		}
		helpers.JsonResponse(ctx, 200, "food fetched successfully", food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var food models.Food
		var menu models.Menu

		if err := ctx.BindJSON(&food); err != nil {
			helpers.NewError(ctx, 400, err.Error())
			return
		}
		validationErr := validate.Struct(food)
		if validationErr != nil {
			helpers.NewError(ctx, 400, validationErr.Error())
			return
		}

		err := menuCollection.FindOne(c, bson.M{"menu_id": food.Menu_Id}).Decode(&menu)
		defer cancel()
		if err != nil {
			helpers.NewError(ctx, 500, "menu was not found")
			return
		}
		food.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		food.ID = primitive.NewObjectID()
		food.Food_Id = food.ID.Hex()
		var num = toFixed(food.Price, 2)
		food.Price = &num

		result, inertErr := foodCollection.InsertOne(c, food)
		if inertErr != nil {
			helpers.NewError(ctx, 500, "food item is not added")
			return
		}
		defer cancel()
		helpers.JsonResponse(ctx, 200, "food inserted successfully", result)
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func round(num float64) int {

}

func toFixed(num *float64, precision int) float64 {

}
