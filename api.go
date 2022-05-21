package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ArsFy/countrycontinent"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"
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
			handleError("SqlErr", err)
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
		handleError("SqlErr", err)
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
		handleError("SqlErr", err)

		if user[0].NodeNum > 0 {
			ip := c.PostForm("ip")
			port, _ := strconv.Atoi(c.PostForm("port"))
			name := c.PostForm("name")
			isp := c.PostForm("isp")
			country := c.PostForm("country")
			continent := c.PostForm("continent")

			if ip != "" && port != 0 && name != "" && isp != "" && country != "" && continent != "" {
				var data []NodeType
				err := Db.Select(&data, "SELECT id FROM node WHERE `ip`=?", ip)
				handleError("SqlErr", err)
				if len(data) != 0 {
					c.JSON(200, gin.H{"status": "exist"})
				} else {
					hostname := GetMD5HashCode([]byte(fmt.Sprint(uid) + ip + fmt.Sprint(port) + "htk746!"))
					key := GetMD5HashCode([]byte(fmt.Sprint(uid) + ip + fmt.Sprint(port) + "kk746?"))
					rows, err := Db.Query(
						"INSERT INTO node (`uid`, `ip`, `port`, `hostname`, `name`, `continent`, `country`, `isp`, `addtime`, `key`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
						uid.(int), ip, port, hostname, name, continent, country, isp, time.Now().Unix(), key,
					)
					handleError("SqlErr", err)
					rows2, err := Db.Query("UPDATE user SET node_num=node_num-1 WHERE `uid`=?", uid.(int))
					handleError("SqlErr", err)
					defer func() {
						rows.Close()
						rows2.Close()
					}()
				}
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

func getIpInfo(c *gin.Context) {
	ip := c.PostForm("ip")
	if ip != "" {
		var data gin.H

		resp, _ := req.C().
			SetTimeout(5 * time.Second).
			R().
			SetResult(&data).
			Get("https://ipinfo.io/" + ip + "?token=" + config["ipinfo_token"].(string))

		if resp.IsSuccess() {
			if data["country"] != nil {
				c.JSON(200, gin.H{
					"country":   data["country"],
					"isp":       data["asn"].(map[string]interface{})["name"],
					"continent": countrycontinent.CountryGetContinent(data["country"].(string)),
					"ip":        data["ip"],
				})
			} else {
				c.JSON(200, gin.H{"status": "error"})
			}
		} else {
			c.JSON(200, gin.H{"status": "error"})
		}
	} else {
		c.JSON(200, gin.H{"status": "empty"})
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
		handleError("SqlErr", err)
		defer rows.Close()
		c.JSON(200, gin.H{"status": "ok"})
	} else {
		c.JSON(200, gin.H{"status": "empty"})
	}
}
