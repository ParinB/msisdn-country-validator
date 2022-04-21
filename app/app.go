package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

var (
	router *gin.Engine
)

func init()  {
	router = gin.Default()
	router.Use(cors.Default())
}

func StartApp()  {
	mapurls()
	if err := router.Run(":3000"); err !=nil {
		panic(err)
	}
}