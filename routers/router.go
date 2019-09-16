package routers

import (
	"wechart/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/wechart/msg/send", &controllers.WechartController{}, "*:Send_Msg")
}
