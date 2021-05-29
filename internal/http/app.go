package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	MapRoutes(router)

	return router
}

func MapRoutes(r *gin.Engine) {
	r.GET("/ping", ping)
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
