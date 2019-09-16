package main

import (
	_ "wechart/routers"
	_ "wechart/models"
	"github.com/astaxie/beego"
	"wechart/cron"
)

func main() {
	go cron.SyncToken()
	beego.Run()
}

