package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getMyNode(c *gin.Context) {
	// Test
	uid := 1

	var data []NodeType
	err := Db.Select(&data, "SELECT * FROM node WHERE `uid`=?", uid)
	if err != nil {
		fmt.Println("SqlErr", err)
	}

	c.JSON(200, gin.H{"status": "ok", "node": data, "len": len(data)})
}
