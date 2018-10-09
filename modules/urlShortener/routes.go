package urlShortener

import "github.com/gin-gonic/gin"

// InitAPIRoutes -
func InitAPIRoutes(router *gin.RouterGroup) {
	router.POST("/urls", postUrlEndpoint)
}

// InitRoutes -
func InitRoutes(router *gin.Engine) {
	router.GET("/:code", getUrlEndpoint)
}
