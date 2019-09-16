package controllers

import (
	"wechart/models"
	"encoding/json"
	"fmt"
)

type WechartController struct {
	BaseController
}

func (c *WechartController) Send_Msg()  {
	var info models.ImMsg
	//fmt.Println(c.Ctx.Input.RequestBody)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"errcode": -1,
			"errmsg": "参数错误",
		}
	c.ServeJSON()
	}
	fmt.Println("info:", info.ToUser, info.Content)
	e := models.MsgSend(info.ToUser, info.ToParty, info.Content)
	if e != nil {
		c.Data["json"] = map[string]interface{}{
			"errcode": -1,
			"errmsg": e.Error(),
		}
		c.ServeJSON()
	}
	c.Data["json"] = map[string]interface{}{
		"errcode": 0,
		"errmsg": "",
	}
	c.ServeJSON()
}