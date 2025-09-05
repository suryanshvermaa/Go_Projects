package controllers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suryanshvermaa/restaurant-management/database"
	"github.com/suryanshvermaa/restaurant-management/helpers"
	"github.com/suryanshvermaa/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

// path: /menus
// method: GET
// description: get all menus
// returns: A JSON response containing the list of menus
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

// path: /menus/:menu_id
// method: GET
// description: get a menu by menu_id
// returns: A JSON response containing the menu details
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

// path: /menus
// method: POST
// description: CreateMenu creates a new menu
// returns: A JSON response containing the created menu details
func CreateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var menu models.Menu
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		if err := ctx.BindJSON(&menu); err != nil {
			helpers.NewError(ctx, 400, err.Error())
			return
		}
		validationErr := validate.Struct(menu)
		if validationErr != nil {
			helpers.NewError(ctx, 400, validationErr.Error())
			return
		}
		menu.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Menu_Id = menu.ID.Hex()
		result, insertErr := menuCollection.InsertOne(c, menu)
		defer cancel()
		if insertErr != nil {
			helpers.NewError(ctx, 500, insertErr.Error())
			return
		}
		helpers.JsonResponse(ctx, 200, "menu created successfully", result)
	}
}

func inTimeSpace(start, end, check time.Time) bool {
	if check.After(start) && check.Before(end) {
		return true
	}
	return false
}

// path: /menus/:menu_id
// method: PATCH
// description: UpdateMenu updates an existing menu by its menu_id
// returns: A JSON response indicating the success or failure of the update operation
func UpdateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		if err := ctx.BindJSON(&menu); err != nil {
			helpers.NewError(ctx, 400, err.Error())
			return
		}

		menuId := ctx.Param("menu_id")
		filter := bson.M{"menu_id": menuId}
		var updateObj bson.D

		if menu.Start_Date != nil && menu.End_Date != nil {
			if !inTimeSpace(*menu.Start_Date, *menu.End_Date, time.Now()) {
				helpers.NewError(ctx, 400, "kindly retype the time")
				return
			}
			updateObj = append(updateObj, bson.E{Key: "start_date", Value: menu.Start_Date})
			updateObj = append(updateObj, bson.E{Key: "end_date", Value: menu.End_Date})
		}
		if menu.Name != "" {
			updateObj = append(updateObj, bson.E{Key: "name", Value: menu.Name})
		}
		if menu.Category != "" {
			updateObj = append(updateObj, bson.E{Key: "category", Value: menu.Category})
		}
		menu.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: menu.Updated_At})

		upsert := true
		opt := options.UpdateOne().SetUpsert(upsert)
		result, err := menuCollection.UpdateOne(c, filter, bson.D{{Key: "$set", Value: updateObj}}, opt)
		defer cancel()
		if err != nil {
			helpers.NewError(ctx, 500, "menu update failed")
			return
		}
		helpers.JsonResponse(ctx, 200, "menu updated successfully", result)
	}
}
