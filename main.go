package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/moos3/skracacz/server"
	"github.com/moos3/skracacz/utils"
	"github.com/vharitonsky/iniflags"
)

func main() {
	utils.Init()
	iniflags.Parse()
	config := utils.GetConfig()

	router := gin.New()
	server.InitRoutes(router)
	router.Run(fmt.Sprintf(":%d", *config.AppPort))

}
