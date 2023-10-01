package auth

import (
	"alienstock-auth-api/config"
	"alienstock-auth-api/models"
	"alienstock-auth-api/util"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description login endpoint
// @Tags Auth
// @Accept json
// @Produce json
// @Param parameter body models.UserModel true "PARAM"
// @Success 200 {string} Register
// @Router /register [post]
func Register(ctx *gin.Context) {
	var user models.UserModel

	// jsonData, err := io.ReadAll(ctx.Request.Body)
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword := util.HashString(user.Password)
	user.Password = hashedPassword

	client, err := config.MongoConfig()
	if err != nil {
		// Handle kesalahan konfigurasi MongoDB
		ctx.JSON(http.StatusInternalServerError, UserResponse{Message: "Databases failed"})
		return
	}
	defer client.Disconnect(context.Background())
	collection := client.Database("AlienStock").Collection("User")

	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
	return

}
