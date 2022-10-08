package routes

import (
	"github.com/bobbyrc/flowmeter/src/routes/api/v1/sample"
	"github.com/bobbyrc/flowmeter/src/routes/api/v1/status"
	"github.com/gin-gonic/gin"
)

func Initialize() *gin.Engine {
	router := gin.Default()

	apiRoutes := router.Group("/api")
	{
		v1Routes := apiRoutes.Group("/v1")
		{
			statusRoutes := v1Routes.Group("/status")
			{
				statusRoutes.GET("/health", status.HealthHandler)
			}
			sampleRoutes := v1Routes.Group("/sample")
			{
				sampleRoutes.POST("", sample.TakeSampleHandler)
			}
		}
	}

	return router
}
