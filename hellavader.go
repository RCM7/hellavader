package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()
	router.Use(cors.Default())
	// Ping test
	router.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "It's alive!")
	})

	router.GET("/hellavader", func(c *gin.Context) {
		c.String(http.StatusOK, "The Force is strong with this one. Vader is commenting on the skills of an X-Wing pilot that is trying to blow up the Death Star, not realizing he’s trying to kill his own son!")
	})

	router.Run()
}
