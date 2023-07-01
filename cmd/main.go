package main

import (
	"lbc/fizzbuzz/controller"
	logger "lbc/fizzbuzz/mdw"
	"lbc/fizzbuzz/stats"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	logger.Init()
	log.Infoln("Server starting")

	//Create singleton
	stats.GetInstance()

	router := gin.Default()
	router.Use(gin.Recovery())

	controller.RegisterEndpoints(router)

	router.Run("0.0.0.0:8080")

	log.Infoln("Server stopped")
}
