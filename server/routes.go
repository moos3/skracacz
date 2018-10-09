package server

import (
	"github.com/gin-gonic/gin"
	"github.com/moos3/skracacz/modules/urlShortener"
)

func InitRoutes(router *gin.Engine) {
	urlShortener.InitRoutes(router)

	api := router.Group("/api")
	{
		urlShortener.InitAPIRoutes(api)
	}
}
