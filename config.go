package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var sqlConf struct {
	MysqlHost        string `json:"mysql_host"`
	MysqlUser        string `json:"mysql_user"`
	MysqlPass        string `json:"mysql_pass"`
	MysqlDatabaseCow string `json:"mysql_database_cow"`
	MysqlDatabaseApi string `json:"mysql_database_api"`
	MysqlPort        string `json:"mysql_port"`
}

var config map[string]interface{}

func init() {
	// Db Config
	dbFile, err := ioutil.ReadFile("./config/mysql.json")
	if err != nil {
		fmt.Println("Config Err", err)
	}
	json.Unmarshal(dbFile, &sqlConf)
	// Other Config
	configFile, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Println("Config Err", err)
	}
	json.Unmarshal(configFile, &config)
}
