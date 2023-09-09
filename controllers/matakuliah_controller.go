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

var MatakuliahCollection *mongo.Collection = configs.GetCollection(configs.DB, "matakuliah")
var validate_Matakuliah = validator.New()

func CreateMatakuliah() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var matakuliah models.Matakuliah
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&matakuliah); err != nil {
			c.JSON(http.StatusBadRequest, responses.MatakuliahResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_Matakuliah.Struct(&matakuliah); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.MatakuliahResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newMatakuliah := models.Matakuliah{
			Nama_matkul:    matakuliah.Nama_matkul,
			SKS:            matakuliah.SKS,
			Dosen_pengampu: matakuliah.Dosen_pengampu,
			Email:          matakuliah.Email,
		}

		result, err := MatakuliahCollection.InsertOne(ctx, newMatakuliah)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.MatakuliahResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.MatakuliahResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetMatakuliah() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		matakuliahID := c.Param("matakuliahGetID")
		var matakuliah models.Matakuliah
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(matakuliahID)

		err := MatakuliahCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&matakuliah)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.MatakuliahResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.MatakuliahResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": matakuliah}})
	}
}

func EditMatakuliah() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		matakuliahID := c.Param("matakuliahID")
		var matakuliah models.Matakuliah
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(matakuliahID)

		//validate the request body
		if err := c.BindJSON(&matakuliah); err != nil {
			c.JSON(http.StatusBadRequest, responses.MatakuliahResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_Matakuliah.Struct(&matakuliah); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.MatakuliahResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"Nama_matkul": matakuliah.Nama_matkul, "SKS": matakuliah.SKS, "Dosen_pengampu": matakuliah.Dosen_pengampu, "Email": matakuliah.Email}
		result, err := MatakuliahCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.MatakuliahResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated LokasiTemuan details
		var updatedLokasiTemuan models.Matakuliah
		if result.MatchedCount == 1 {
			err := MatakuliahCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedLokasiTemuan)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.MatakuliahResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.MatakuliahResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedLokasiTemuan}})
	}
}

func DeleteMatakuliah() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		matakuliahID := c.Param("matakuliahID")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(matakuliahID)

		result, err := MatakuliahCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.MatakuliahResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.MatakuliahResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "matakuliah with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.MatakuliahResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "matakuliah successfully deleted!"}},
		)
	}
}

func GetAllMatakuliahs() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var matakuliahs []models.Matakuliah
		defer cancel()

		results, err := MatakuliahCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.MatakuliahResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleLokasiTemuan models.Matakuliah
			if err = results.Decode(&singleLokasiTemuan); err != nil {
				c.JSON(http.StatusInternalServerError, responses.MatakuliahResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			matakuliahs = append(matakuliahs, singleLokasiTemuan)
		}

		c.JSON(http.StatusOK,
			responses.MatakuliahResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": matakuliahs}},
		)
	}
}
