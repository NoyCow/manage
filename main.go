package main

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// HTTP
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("server", "NoyPL 1.0")
		c.Writer.Header().Set("Content-Type", "application/json;charset=utf-8")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	store, _ := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("COW_SESSION", store))
	router.Use(cors())
	router.Use(gin.Recovery())
	// 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": 404})
	})
	// Status
	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": 200})
	})

	// API
	router.POST("/api/login", login)
	router.POST("/api/getmynode", getMyNode)
	router.POST("/api/update", update) // Update Node Date
	router.POST("/api/addnode", addNode)

	// Run
	fmt.Println("NoyPL Starting ...")
	router.Run(":80")
}
