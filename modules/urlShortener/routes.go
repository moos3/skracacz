package urlShortener

import "github.com/gin-gonic/gin"

func InitApiRoutes(router *gin.RouterGroup) {
	router.POST("/urls", postUrlEndpoint)
}

func InitRoutes(router *gin.Engine) {
	router.GET("/:code", getUrlEndpoint)
}
