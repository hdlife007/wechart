package models

import (
	"wechart/conf"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"time"
	"fmt"
	"errors"
)

//定义消息结构
type ImMsg struct {
	ToUser	string	`json:"touser"`
	ToParty	string	`json:"toparty"`
	MsgType	string	`json:"msgtype"`
	AgentId	int64	`json:"agentid"`
	Content	string	`json:"content"`
}

type SendMsgReq struct {
	conf.WechartResult
	Invaliduser	string	`json:"invaliduser"`
	Invalidparty	string	`json:"invalidparty"`
}

//调用企业微信接口发送消息
func MsgSend(touser, toparty, content string) error  {
	token := conf.SafeTokenRet.Get()
	agentid := beego.AppConfig.String("agent_id")
	fmt.Println("token:", token, "agentid:", agentid)
	msgtype := "text"
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token)

	req := httplib.Post(url)
	text := make(map[string]string)
	text["content"] = content
	req.JSONBody(map[string]interface{}{"touser":touser, "toparty":toparty, "msgtype":msgtype, "agentid":agentid, "text":text})
	req.SetTimeout(1*time.Second, 3*time.Second)

	var rep SendMsgReq

	ok := req.ToJSON(&rep)
	if ok != nil {
		return ok
	}
	errcode := rep.Errcode
	if errcode !=0 {
		return errors.New(rep.Errmsg)
	}
	return nil
}
