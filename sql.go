package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", sqlConf.MysqlUser+":"+sqlConf.MysqlPass+"@tcp("+sqlConf.MysqlHost+":"+sqlConf.MysqlPort+")/"+sqlConf.MysqlDatabase+"?charset=utf8mb4")
	if err != nil {
		fmt.Println("DbErr", err)
		return
	}
	database.SetMaxOpenConns(3000)
	database.SetMaxIdleConns(500)

	Db = database

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
	Area      string `db:"area" json:"area"`
	Img       int    `db:"img" json:"img"`
	Frequency int    `db:"frequency" json:"frequency"`
	Broadband int    `db:"broadband" json:"broadband"`
	Status    int    `db:"status" json:"status"`
	LastTime  int    `db:"lasttime" json:"lasttime"`
	AddTime   int    `db:"addtime" json:"addtime"`
}
