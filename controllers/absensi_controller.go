package controllers

import (
	"context"
	"gin-mongo-api/configs"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var AbsensiCollection *mongo.Collection = configs.GetCollection(configs.DB, "Abensi")
var validate_absensi = validator.New()

func CreateAbsensi() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var absensi models.Absensi
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&absensi); err != nil {
			c.JSON(http.StatusBadRequest, responses.AbsensiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_absensi.Struct(&absensi); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AbsensiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newAbsensi := models.Absensi{
			Id:      primitive.NewObjectID(),
			Nama_mk: absensi.Nama_mk,
			Tanggal: absensi.Tanggal,
			Checkin: absensi.Checkin,
		}

		result, err := AbsensiCollection.InsertOne(ctx, newAbsensi)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AbsensiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.AbsensiResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetAabsensi() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("absensiGetId")
		var user models.Absensi
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := AbsensiCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AbsensiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AbsensiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}

func EditAbsensi() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		absensiId := c.Param("absensiID")
		var absensi models.Absensi
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(absensiId)

		//validate the request body
		if err := c.BindJSON(&absensi); err != nil {
			c.JSON(http.StatusBadRequest, responses.AbsensiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_absensi.Struct(&absensi); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AbsensiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"name": absensi.Nama_mk, "location": absensi.Tanggal, "title": absensi.Checkin}
		result, err := AbsensiCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AbsensiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated user details
		var updatedAbsensi models.Absensi
		if result.MatchedCount == 1 {
			err := AbsensiCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedAbsensi)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.AbsensiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.AbsensiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedAbsensi}})
	}
}

func DeleteAabsensi() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		absensiId := c.Param("absensiID")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(absensiId)

		result, err := AbsensiCollection.DeleteOne(ctx, bson.M{"id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AbsensiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.AbsensiResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Absensi with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.AbsensiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Absensi successfully deleted!"}},
		)
	}
}

func GetAllAbssenis() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var absensis []models.Absensi
		defer cancel()

		results, err := AbsensiCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AbsensiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser models.Absensi
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, responses.AbsensiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			absensis = append(absensis, singleUser)
		}

		c.JSON(http.StatusOK,
			responses.AbsensiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": absensis}},
		)
	}
}
