package main

import (
	"time"

	"github.com/robfig/cron/v3"
)

func stopLast() {
	c := cron.New()
	c.AddFunc("0 */1 * * * ?", func() {
		Db.Query("UPDATE node SET `status` = 0 WHERE `lasttime` < ?", time.Now().Unix()-1800)
	})
	c.Start()
}
