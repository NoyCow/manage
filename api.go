package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	user := c.PostForm("user")
	pass := c.PostForm("pass")

	if user != "" && pass != "" {
		var data []UserType
		DbApi.Select(&data, "SELECT uid, username, email FROM user WHERE (`username`=? or `email`=?) and `password`=?", user, user, GetSHA256HashCode([]byte(pass)))
		if len(data) != 0 {
			// Add User
			rows, err := Db.Query("INSERT IGNORE INTO user (`uid`) VALUES(?)", data[0].Uid)
			if err != nil {
				fmt.Println("SqlErr", err)
			}
			rows.Close()
			// Save Session
			session := sessions.Default(c)
			session.Set("uid", data[0].Uid)
			session.Save()
			c.JSON(200, gin.H{"status": "ok"})
		} else {
			c.JSON(200, gin.H{"status": "user_pass_err"})
		}
	} else {
		c.JSON(200, gin.H{"status": "empty"})
	}
}

func getMyNode(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("uid")
	if uid != nil {
		var data []NodeType
		err := Db.Select(&data, "SELECT * FROM node WHERE `uid`=?", uid.(int))
		if err != nil {
			fmt.Println("SqlErr", err)
		}
		c.JSON(200, gin.H{"status": "ok", "node": data, "len": len(data), "uid": uid.(int)})
	} else {
		c.JSON(200, gin.H{"status": "login"})
	}
}

func addNode(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("uid")
	if uid != nil {
		var user []CowUserType
		err := Db.Select(&user, "SELECT uid, node_num FROM user WHERE `uid`=?", uid.(int))
		if err != nil {
			fmt.Println("SqlErr", err)
		}

		if user[0].NodeNum > 0 {
			ip := c.PostForm("ip")
			port, _ := strconv.Atoi(c.PostForm("port"))
			name := c.PostForm("name")

			if ip != "" && port != 0 && name != "" {
				// Add Code
			} else {
				c.JSON(200, gin.H{"status": "empty"})
			}
		} else {
			c.JSON(200, gin.H{"status": "quota"})
		}
	} else {
		c.JSON(200, gin.H{"status": "login"})
	}
}

func update(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	key := c.PostForm("key")
	img, _ := strconv.Atoi(c.PostForm("img"))
	frequency, _ := strconv.Atoi(c.PostForm("frequency"))
	broadband, _ := strconv.Atoi(c.PostForm("broadband"))

	if key != "" && id != 0 {
		rows, err := Db.Query("UPDATE node SET `img`=? , `frequency`=? , `broadband`=?, `time`=?, `status`=1 WHERE `id`=? and key=?", img, frequency, broadband, time.Now().Unix(), id, key)
		if err != nil {
			fmt.Println("SqlError", err)
		}
		rows.Close()
		c.JSON(200, gin.H{"status": "ok"})
	} else {
		c.JSON(200, gin.H{"status": "empty"})
	}
}
