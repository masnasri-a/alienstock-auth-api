package profile

import (
	"alienstock-auth-api/config"
	"alienstock-auth-api/models"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserResponse struct {
	Message string `json:"message"`
}

// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description login endpoint
// @Tags Profile
// @Accept json
// @Produce json
// @Param        fullname    query     string  false  "data search by fullname"
// @Success 200 {string} GetProfile
// @Router /getProfile [get]
func GetSelfProfile(ctx *gin.Context) {
	fmt.Println(ctx)
	fullname := ctx.Query("fullname")
	fmt.Printf("Fullname : %s\n", fullname)
	var user models.UserModel
	var responseFind models.ResponseUserModel
	client, err := config.MongoConfig()
	if err != nil {
		// Handle kesalahan konfigurasi MongoDB
		// panic(err)
		ctx.JSON(http.StatusInternalServerError, UserResponse{Message: "Databases failed"})
		return
	}
	defer client.Disconnect(context.Background())
	collection := client.Database("AlienStock").Collection("User")
	filter := bson.D{{Key: "fullname", Value: user.FullName}}
	err = collection.FindOne(context.Background(), filter).Decode(&responseFind)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			ctx.JSON(http.StatusUnauthorized, UserResponse{Message: "Authentication failed"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, UserResponse{Message: "Databases failed"})
			return
		}
	}
	ctx.JSON(http.StatusCreated, responseFind)
	return

}
