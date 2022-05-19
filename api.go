package main

import (
	"fmt"
	"strconv"
	"time"

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

func update(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	key := c.PostForm("key")
	img, _ := strconv.Atoi(c.PostForm("img"))
	frequency, _ := strconv.Atoi(c.PostForm("frequency"))
	broadband, _ := strconv.Atoi(c.PostForm("broadband"))

	if key != "" && id != 0 {
		_, err := Db.Query("UPDATE node SET `img`=? , `frequency`=? , `broadband`=?, `time`=?, `status`=1 WHERE `id`=? and key=?", img, frequency, broadband, time.Now().Unix(), id, key)
		if err != nil {
			fmt.Println("SqlError", err)
		}
		c.JSON(200, gin.H{"status": "ok"})
	} else {
		c.JSON(200, gin.H{"status": "empty"})
	}
}
