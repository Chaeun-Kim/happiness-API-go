package main

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

func main() {
	router := InitRouter()
	slog.Info("Starting up application at port 5000...")
	router.Run(":5000")
}

func InitRouter() *gin.Engine {
	router := gin.Default()

	data := InitHappinessIndexData()
	router.GET("/ping", PingHandler())
	router.GET("/happiness/:facet/:facet_id", HappinessByFacetIdHandler(data))
	router.POST("/happiness/:facet", HappinessByFacetIdsHandler(data))

	return router
}
