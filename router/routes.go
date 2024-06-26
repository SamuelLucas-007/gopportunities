package router

import (
	"github.com/gin-gonic/gin"

	"github.com/SamuelLucas-007/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.GetOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.PATCH("/opening", handler.PatchOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}
}
