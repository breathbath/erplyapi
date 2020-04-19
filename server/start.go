package server

import (
	"github.com/breathbath/erplyapi/erply"
	"github.com/breathbath/go_utils/utils/env"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

//Start main entry point for the gin server
func Start() error {
	router := gin.Default()
	
	log := logrus.New()
	
	router.Use(ginlogrus.Logger(log), gin.Recovery())

	customerRoute := router.Group("/customers")
	{
		customerRoute.GET("/:ids", erply.GetCustomersByIdHandler)
		customerRoute.POST("/", erply.CreateCustomerHandler)
	}

	addr := ":" + env.ReadEnv("REST_SERVER_PORT", "8082")

	return router.Run(addr)
}
