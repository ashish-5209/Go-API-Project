package main

import (
	"log"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

var (
	mainRouter *gin.Engine
)

func init() {
	mainRouter = gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(mainRouter)

	mainRouter.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	AddHeapRouter(mainRouter)
}

func main() {
	err := mainRouter.Run(":3000")
	if err != nil {
		log.Println("Failed to start server:", err)
	}
	log.Println("Server started")
}
