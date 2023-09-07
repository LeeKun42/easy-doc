package job

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

var c *cron.Cron

func init() {
	c = cron.New(cron.WithSeconds())
	// go cron 规则参考 https://pkg.go.dev/github.com/robfig/cron?utm_source=godoc
	//添加执行任务
	c.AddFunc("*/30 * * * * *", func() {
		fmt.Println("job exec")
	}) //每日9点执行一次
}

func Start() {
	c.Start()
}
