package server

import (
	"github.com/breathbath/erplyapi/auth"
	"github.com/breathbath/erplyapi/db"
	"github.com/breathbath/erplyapi/metrics"
	"github.com/breathbath/erplyapi/reports"
	"github.com/breathbath/go_utils/utils/env"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/thinkerou/favicon"
	ginlogrus "github.com/toorop/gin-logrus"
)

//Start main entry point for the gin server
func Start() error {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.Use(favicon.New("./favicon.ico"))
	router.Use(ginlogrus.Logger(logrus.StandardLogger()))
	router.Use(gin.Recovery())

	frontAuthMiddleWare, err := auth.BuildFrontMiddleWare()
	if err != nil {
		return err
	}
	router.POST("/front-login", frontAuthMiddleWare.LoginHandler)
	router.POST("/front-refresh", frontAuthMiddleWare.RefreshHandler)

	backAuthMiddleWare, err := auth.BuildBackMiddleWare()
	if err != nil {
		return err
	}
	router.POST("/back-login", backAuthMiddleWare.LoginHandler)
	router.POST("/back-refresh", backAuthMiddleWare.RefreshHandler)

	baseDB, err := db.NewDb()
	if err != nil {
		return err
	}
	visitsDb := db.Visits{Db: baseDB}
	visitsEndpoint := metrics.Endpoint{VisitStore: visitsDb}

	visitsRoute := router.Group("/visits")
	visitsRoute.Use(frontAuthMiddleWare.MiddlewareFunc())
	{
		visitsRoute.POST("", visitsEndpoint.CreateVisitsHandler)
	}

	visitStatsHandler := reports.ReportsHandler{ReportsProvider: visitsDb}
	graphsRoute := router.Group("/reports")
	graphsRoute.Use(backAuthMiddleWare.MiddlewareFunc())
	{
		graphsRoute.GET("visits-by-hour.:format", visitStatsHandler.VisitsByHourHandler)
		graphsRoute.GET("visits-by-location.:format", visitStatsHandler.VisitsByLocationHandler)
		graphsRoute.GET("visits-by-day.:format", visitStatsHandler.VisitsByDayHandler)
		graphsRoute.GET("visits-by-month.:format", visitStatsHandler.VisitsByMonthHandler)
	}

	docsPath := env.ReadEnv("DOCS_PATH", "")
	if docsPath != "" {
		router.Static("/docs", docsPath)
	}

	addr := ":" + env.ReadEnv("REST_SERVER_PORT", "8082")

	return router.Run(addr)
}
