package handler

import "github.com/gin-gonic/gin"

func PatchOpeningHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "PATCH /api/v1/opening",
	})
}
