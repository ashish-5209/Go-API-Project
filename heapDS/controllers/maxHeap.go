package controllers

import (
	"go-api/heapDS"
	"go-api/heapDS/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuildMaxHeap(ctx *gin.Context) {
	req := &models.MaxHeap{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid Request",
		})
		return
	}

	if req.List == nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid Request",
		})
		return
	}

	obj := &heapDS.MaxHeap{}

	obj.BuildMaxHeap(ctx, req.List)

	if len(obj.MaxHeapList) == 0 {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"success": false,
			"message": "List is empty",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Max Heap List",
		"data":    obj.MaxHeapList,
	})

}

func ExtractMax(ctx *gin.Context) {
	req := &models.MaxHeap{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid Request",
		})
		return
	}

	if req.List == nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid Request",
		})
		return
	}

	obj := &heapDS.MaxHeap{}
	obj.MaxHeapList = make([]int, len(req.List))
	// Copy elements from req.List to heap.MaxHeapList
	copy(obj.MaxHeapList, req.List)
	response := obj.ExtractMax(ctx, len(req.List))

	if response == -1 {
		ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"success": false,
			"message": "List is empty",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Max Heap List",
		"data":    response,
	})

}
