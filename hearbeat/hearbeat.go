package hearbeat

import (
	"myStorage/rabbitmq"
	"os"
	"time"
)

//数据服务节点发送心跳信号 5秒间隔
func StartHearbeat()  {
	//根据传入的当前节点数据服务需要启动的服务地址创建一个数据服务对象
	q := rabbitmq.RNew(os.Getenv("RABBITMQ_SERVER"))
	defer q.RClose()
	for  {
		//发送的是数据服务的监听地址还是接口服务的监听地址？？？
		q.RPublish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(5 * time.Second)
	}
}
