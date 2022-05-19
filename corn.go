package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func stopLast() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 */5 * * * ?", func() {
		_, err := Db.Query("UPDATE node SET `status` = 0 WHERE `lasttime` < ?", time.Now().Unix()-1800)
		if err != nil {
			fmt.Println("SqlError", err)
		}
	})
	c.Start()
}
