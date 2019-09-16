# wechart
微信告警接口
依赖beego，运行不成功的话先go get github.com/beego 安装beego
conf目录下app.conf中的corpid（企业ID）， corpsecret（自建应用的Secret）， agent_id（自建应用的AgentId） 替换为自己的企业微信相关信息
调用企业微信接口发送消息格式样例（ToUser和ToParty不可同时为空）：
curl -i -X POST -H "'Content-type':'application/json'" -d '{"ToUser":"", "ToParty":"", "Content":""}' http://ip:8088/wechart/msg/send/
仅实现了发送文本消息，其他格式待扩展
