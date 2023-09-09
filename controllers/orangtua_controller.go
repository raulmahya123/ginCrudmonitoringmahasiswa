package controllers

import (
	"context"
	"gin-mongo-api/configs"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var orangTuaCollection *mongo.Collection = configs.GetCollection(configs.DB, "orangtua")
var validate_orangtua = validator.New()

func CreateOrangtua() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var orangtua models.OrangTua
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&orangtua); err != nil {
			c.JSON(http.StatusBadRequest, responses.OrangtuaResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_orangtua.Struct(&orangtua); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.OrangtuaResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newOrangtua := models.OrangTua{
			Id:           primitive.NewObjectID(),
			Nama_ortu:    orangtua.Nama_ortu,
			Phone_number: orangtua.Phone_number,
			Email:        orangtua.Email,
		}

		result, err := orangTuaCollection.InsertOne(ctx, newOrangtua)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.OrangtuaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.OrangtuaResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetOrangtua() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		orangtuaID := c.Param("orangtuaGetID")
		log.Println(orangtuaID)
		var orangtua models.OrangTua
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(orangtuaID)

		err := orangTuaCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&orangtua)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.OrangtuaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.OrangtuaResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": orangtua}})
	}
}

func EditOrangtua() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		orangtuaID := c.Param("orangtuaID")
		var orangtua models.OrangTua
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(orangtuaID)

		//validate the request body
		if err := c.BindJSON(&orangtua); err != nil {
			c.JSON(http.StatusBadRequest, responses.OrangtuaResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_orangtua.Struct(&orangtua); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.OrangtuaResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"Nama_ortu": orangtua.Nama_ortu, "Phone_number": orangtua.Phone_number, "Email": orangtua.Email}
		result, err := orangTuaCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.OrangtuaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated invertebrata details
		var updatedInvertebrata models.OrangTua
		if result.MatchedCount == 1 {
			err := orangTuaCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedInvertebrata)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.OrangtuaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.OrangtuaResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedInvertebrata}})
	}
}

func DeleteOrangtua() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		orangtuaID := c.Param("orangtuaID")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(orangtuaID)

		result, err := orangTuaCollection.DeleteOne(ctx, bson.M{"id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.OrangtuaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.OrangtuaResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "orangtua with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.OrangtuaResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "orangtua successfully deleted!"}},
		)
	}
}

func GetAllOrangtuas() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var invertebratas []models.OrangTua
		defer cancel()

		results, err := orangTuaCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.OrangtuaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleInvertebrata models.OrangTua
			if err = results.Decode(&singleInvertebrata); err != nil {
				c.JSON(http.StatusInternalServerError, responses.OrangtuaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			invertebratas = append(invertebratas, singleInvertebrata)
		}

		c.JSON(http.StatusOK,
			responses.OrangtuaResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": invertebratas}},
		)
	}
}
