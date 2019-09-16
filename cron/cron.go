package cron

import (
	"wechart/conf"
	"github.com/astaxie/beego/httplib"
	"time"
	"github.com/astaxie/beego"
	"fmt"
	"errors"
	"wechart/logs"
)

type TokenReq struct {
	conf.WechartResult
	AccessToken	string	`json:"access_token"`
}

func getToken(secret string) (string, error)  {
	corpid := beego.AppConfig.String("corpid")

	//fmt.Println("secret:", secret, "corpid:", corpid)

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", corpid, secret)
	req := httplib.Get(url)

	//fmt.Println(req)

	req.SetTimeout(1*time.Second, 3*time.Second)
	var ret TokenReq
	ok := req.ToJSON(&ret)
	if ok != nil {
		return "", ok
	}

	errCode := ret.Errcode
	if errCode != 0 {
		return "", errors.New(ret.Errmsg)
	}

	return ret.AccessToken, nil

}

func syncToken()  {
	secret := beego.AppConfig.String("corpsecret")

	token, err := getToken(secret)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("token:", token)
	conf.SafeTokenRet.ReInit(token)
}

func SyncToken()  {
	interval := beego.AppConfig.DefaultInt64("access_token_interval", 7200)
	duration := time.Duration(interval) * time.Second
	for {
		currentTime := time.Now().Format("2006-01-02 15:04:05")

		logs.Log.Info("cron_sync_token_time: %v", currentTime)
		syncToken()
		time.Sleep(duration)
	}
}
