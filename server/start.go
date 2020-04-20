package server

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/breathbath/erplyapi/auth"
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
	
	router.Use(ginlogrus.Logger(log))
	router.Use(gin.Recovery())

	authMiddleware, err := auth.BuildMiddleWare()
	if err != nil {
		return err
	}
	router.POST("/login", authMiddleware.LoginHandler)

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router.POST("/refresh", authMiddleware.RefreshHandler)

	customerRoute := router.Group("/customers")
	customerRoute.Use(authMiddleware.MiddlewareFunc())
	{
		customerRoute.GET("/:ids", erply.GetCustomersByIdHandler)
		customerRoute.POST("/", erply.CreateCustomerHandler)
	}

	docsPath := env.ReadEnv("DOCS_PATH", "")
	if docsPath != "" {
		router.Static("/docs", docsPath)
	}

	addr := ":" + env.ReadEnv("REST_SERVER_PORT", "8082")

	return router.Run(addr)
}
