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

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			helpers.NewError(ctx, 400, err.Error())
			return
		}
		var allMenus []bson.M
		if err = result.All(c, &allMenus); err != nil {
			helpers.NewError(ctx, 400, err.Error())
			return
		}
		helpers.JsonResponse(ctx, 200, "menus fetched successfully", allMenus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		menuId := ctx.Param("menu_id")
		var menu models.Menu
		err := foodCollection.FindOne(c, bson.M{"menu_id": menuId}).Decode(&menu)
		defer cancel()
		if err != nil {
			helpers.NewError(ctx, 400, err.Error())
			return
		}
		helpers.JsonResponse(ctx, 200, "menu fetched successfully", menu)
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
