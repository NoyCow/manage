package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB
var DbApi *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Open("mysql", sqlConf.MysqlUser+":"+sqlConf.MysqlPass+"@tcp("+sqlConf.MysqlHost+":"+sqlConf.MysqlPort+")/"+sqlConf.MysqlDatabaseCow+"?charset=utf8mb4")
	if err != nil {
		fmt.Println("DbErr", err)
		return
	}
	Db.SetMaxOpenConns(3000)
	Db.SetMaxIdleConns(500)

	DbApi, err = sqlx.Open("mysql", sqlConf.MysqlUser+":"+sqlConf.MysqlPass+"@tcp("+sqlConf.MysqlHost+":"+sqlConf.MysqlPort+")/"+sqlConf.MysqlDatabaseApi+"?charset=utf8mb4")
	if err != nil {
		fmt.Println("DbErr", err)
		return
	}
	DbApi.SetMaxOpenConns(1000)
	DbApi.SetMaxIdleConns(200)

	// Corn
	stopLast()
}

type NodeType struct {
	Id        int    `db:"id" json:"id"`
	Uid       int    `db:"uid" json:"uid"`
	Ip        string `db:"ip" json:"ip"`
	Port      int    `db:"port" json:"port"`
	Hostname  string `db:"hostname" json:"hostname"`
	Name      string `db:"name" json:"name"`
	Continent string `db:"continent" json:"continent"`
	Country   string `db:"country" json:"country"`
	Isp       string `db:"ips" json:"isp"`
	Img       int    `db:"img" json:"img"`
	Frequency int    `db:"frequency" json:"frequency"`
	Broadband int    `db:"broadband" json:"broadband"`
	Status    int    `db:"status" json:"status"`
	LastTime  int    `db:"lasttime" json:"lasttime"`
	AddTime   int    `db:"addtime" json:"addtime"`
	Key       string `db:"key" json:"key"`
}

type UserType struct {
	Uid      int    `db:"uid" json:"uid"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
}

type CowUserType struct {
	Uid     int `db:"uid"`
	NodeNum int `db:"node_num"`
}

type BookType struct {
	Bid int `db:"bid" json:"bid"`
}
