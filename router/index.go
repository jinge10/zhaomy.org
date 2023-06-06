package router

import (
	"github.com/gin-gonic/gin"
)

// 全局 Gin Engine 对象
var router *gin.Engine

func init() {
	router = gin.Default()
	router.Static("assets", "static/assets")
	router.StaticFile("vite.svg", "static/vite.svg")
	router.LoadHTMLGlob("static/index.html")
}

// @return GIN RouterEngine
func GetRouter() *gin.Engine {
	return router
}
