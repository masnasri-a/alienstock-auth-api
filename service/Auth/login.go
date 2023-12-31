package auth

import (
	"alienstock-auth-api/config"
	"alienstock-auth-api/models"
	"alienstock-auth-api/util"
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
// @Tags Auth
// @Accept json
// @Produce json
// @Param parameter body models.LoginModel true "PARAM"
// @Success 200 {string} Login
// @Router /login [post]
func Login(ctx *gin.Context) {
	var user models.LoginModel
	var responseFind models.ResponseUserModel
	// jsonData, err := io.ReadAll(ctx.Request.Body)
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword := util.HashString(user.Password)
	user.Password = hashedPassword
	// Simpan pengguna yang didaftarkan
	client, err := config.MongoConfig()
	if err != nil {
		// Handle kesalahan konfigurasi MongoDB
		ctx.JSON(http.StatusInternalServerError, UserResponse{Message: "Cant connect database"})
		return
	}
	defer client.Disconnect(context.Background())
	collection := client.Database("AlienStock").Collection("User")
	filter := bson.D{{Key: "email", Value: user.Email}, {Key: "password", Value: user.Password}}
	err = collection.FindOne(context.Background(), filter).Decode(&responseFind)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			ctx.JSON(http.StatusUnauthorized, UserResponse{Message: "Authentication failed"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, UserResponse{Message: "Database connection has broken"})
		return
	}

	fmt.Printf("Username : %s\n", responseFind.FullName)
	// Kirim respons
	ctx.JSON(http.StatusCreated, UserResponse{Message: "User logged in successfully"})
	return
}
