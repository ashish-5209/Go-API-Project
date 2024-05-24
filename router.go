package main

import (
	"go-api/heapDS/controllers"

	"github.com/gin-gonic/gin"
)

func AddHeapRouter(router *gin.Engine) {
	heapRouter := router.Group("heap")
	maxHeapRouter := heapRouter.Group("max")
	maxHeapRouter.POST("build", controllers.BuildMaxHeap)
	maxHeapRouter.POST("extract", controllers.ExtractMax)
}
