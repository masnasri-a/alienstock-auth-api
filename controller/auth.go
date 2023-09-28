package controller

import (
	auth "alienstock-auth-api/service/Auth"

	"github.com/gin-gonic/gin"
)

func AuthRouter(route *gin.RouterGroup) {
	route.POST("/login", auth.Login)
	route.POST("/register", auth.Register)
}
