package controller

import (
	profile "alienstock-auth-api/service/Profile"

	"github.com/gin-gonic/gin"
)

func ProfileRouter(route *gin.RouterGroup) {
	route.GET("/getProfile", profile.GetSelfProfile)
}
