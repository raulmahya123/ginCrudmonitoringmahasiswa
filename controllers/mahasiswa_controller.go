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

var MahasiswaCollection *mongo.Collection = configs.GetCollection(configs.DB, "mahasiswa")
var validate_mahasiswa = validator.New()

func CreateMahasiswa() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var mahasiswa models.Mahasiswa
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&mahasiswa); err != nil {
			c.JSON(http.StatusBadRequest, responses.MahasiswaResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_mahasiswa.Struct(&mahasiswa); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.MahasiswaResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newMahasiswa := models.Mahasiswa{
			Id:       primitive.NewObjectID(),
			Nama_mhs: mahasiswa.Nama_mhs,
			NPM:      mahasiswa.NPM,
			Jurusan:  mahasiswa.Jurusan,
			Email:    mahasiswa.Email,
		}

		result, err := MahasiswaCollection.InsertOne(ctx, newMahasiswa)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.MahasiswaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.MahasiswaResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetMahasiswa() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		mahasiswaId := c.Param("mahasiswaGetid")
		var mahasiswa models.Mahasiswa
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(mahasiswaId)

		err := MahasiswaCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&mahasiswa)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.MahasiswaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.MahasiswaResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": mahasiswa}})
	}
}

func EditMahasiswa() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		mahasiswaId := c.Param("mahasiswaID")
		var mahasiswa models.Mahasiswa
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(mahasiswaId)

		//validate the request body
		if err := c.BindJSON(&mahasiswa); err != nil {
			c.JSON(http.StatusBadRequest, responses.MahasiswaResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_mahasiswa.Struct(&mahasiswa); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.MahasiswaResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"Nama_mhs": mahasiswa.Nama_mhs, "NPM": mahasiswa.NPM, "Jurusan": mahasiswa.Jurusan, "Email": mahasiswa.Email}
		result, err := MahasiswaCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.MahasiswaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated Fosil details
		var UpdatedMahasiswa models.Mahasiswa
		if result.MatchedCount == 1 {
			err := MahasiswaCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&UpdatedMahasiswa)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.MahasiswaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.MahasiswaResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": UpdatedMahasiswa}})
	}
}

func DeleteMahasiswa() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		mahasiswaId := c.Param("mahasiswaID")
		log.Println(mahasiswaId)
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(mahasiswaId)

		result, err := MahasiswaCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.MahasiswaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.MahasiswaResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Mahasiswa with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.MahasiswaResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Mahasiswa successfully deleted!"}},
		)
	}
}

func GetAllMahasiswas() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var mahasiswas []models.Mahasiswa
		defer cancel()

		results, err := MahasiswaCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.MahasiswaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleFosil models.Mahasiswa
			if err = results.Decode(&singleFosil); err != nil {
				c.JSON(http.StatusInternalServerError, responses.MahasiswaResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			mahasiswas = append(mahasiswas, singleFosil)
		}

		c.JSON(http.StatusOK,
			responses.MahasiswaResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": mahasiswas}},
		)
	}
}
