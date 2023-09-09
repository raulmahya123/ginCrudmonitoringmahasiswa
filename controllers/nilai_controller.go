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

var nilaiCollection *mongo.Collection = configs.GetCollection(configs.DB, "nilai")
var validate_nilai = validator.New()

func CreateNilai() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var nilai models.Nilai
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&nilai); err != nil {
			c.JSON(http.StatusBadRequest, responses.NilaiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_nilai.Struct(&nilai); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.NilaiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newAbsensi := models.Nilai{
			Id:           primitive.NewObjectID(),
			NPM_ms:       nilai.NPM_ms,
			Presensi:     nilai.Presensi,
			Nilai_akhir:  nilai.Nilai_akhir,
			Grade:        nilai.Grade,
			Tahun_ajaran: nilai.Tahun_ajaran,
		}

		result, err := nilaiCollection.InsertOne(ctx, newAbsensi)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.NilaiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.NilaiResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetNilai() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		NilaiID := c.Param("nialiGetID")
		var nilai models.Nilai
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(NilaiID)

		err := nilaiCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&nilai)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.NilaiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.NilaiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": nilai}})
	}
}

func EditNilai() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		nilaiID := c.Param("nialiID")
		var nilai models.Nilai
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(nilaiID)

		//validate the request body
		if err := c.BindJSON(&nilai); err != nil {
			c.JSON(http.StatusBadRequest, responses.NilaiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_nilai.Struct(&nilai); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.NilaiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"NPM_ms": nilai.NPM_ms, "Presensi": nilai.Presensi, "Nilai_akhir": nilai.Nilai_akhir, "Grade": nilai.Grade, "tahun ajaran": nilai.Tahun_ajaran}
		result, err := nilaiCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.NilaiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated Vertebrata details
		var updatedNilai models.Nilai
		if result.MatchedCount == 1 {
			err := nilaiCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedNilai)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.NilaiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.NilaiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedNilai}})
	}
}

func DeleteNilai() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		DeleteID := c.Param("nialiID")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(DeleteID)

		result, err := nilaiCollection.DeleteOne(ctx, bson.M{"id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.NilaiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.NilaiResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "nilai with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.NilaiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "nilai successfully deleted!"}},
		)
	}
}

func GetAllNilais() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var vertebratas []models.Nilai
		defer cancel()

		results, err := nilaiCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.NilaiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleVertebrata models.Nilai
			if err = results.Decode(&singleVertebrata); err != nil {
				c.JSON(http.StatusInternalServerError, responses.NilaiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			vertebratas = append(vertebratas, singleVertebrata)
		}

		c.JSON(http.StatusOK,
			responses.NilaiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": vertebratas}},
		)
	}
}
