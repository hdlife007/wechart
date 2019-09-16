package conf

import "sync"

type WechartResult struct {
	Errcode	int64	`json:"errcode"`
	Errmsg	string	`json:"errmsg"`
}

type SafeToken struct {
	sync.RWMutex
	Token	string
}

var SafeTokenRet = &SafeToken{}

func (this *SafeToken) ReInit(token string)  {
	this.Lock()
	defer this.Unlock()
	this.Token = token
}

func (this *SafeToken) Get() string  {
	this.Lock()
	defer this.Unlock()
	return this.Token
}
