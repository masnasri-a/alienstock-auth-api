package auth

import (
	"alienstock-auth-api/config"
	"alienstock-auth-api/models"
	"alienstock-auth-api/util"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users []models.UserModel

type UserResponse struct {
	Message string `json:"message"`
}

// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description login endpoint
// @Tags example
// @Accept json
// @Produce json
// @Param parameter body models.LoginModel true "PARAM"
// @Success 200 {string} Login
// @Router /login [post]
func Login(ctx *gin.Context) {
	var user models.UserModel

	// jsonData, err := io.ReadAll(ctx.Request.Body)
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan pengguna yang didaftarkan

	fmt.Printf("Username : %s\n", user.FullName)
	// Kirim respons
	ctx.JSON(http.StatusCreated, UserResponse{Message: "User registered successfully"})

}

// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description login endpoint
// @Tags example
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
		panic(err)
	}
	defer client.Disconnect(context.Background())
	collection := client.Database("AlienStock").Collection("User")

	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})

}
