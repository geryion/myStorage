package main

import (
	"fmt"
	"log"
	"myStorage/hearbeat"
	"myStorage/locate"
	"myStorage/objects"
	"net/http"
	"os"
)

func main()  {
	//数据服务节点心跳检测
	go hearbeat.StartHearbeat()

	//实际定位数据对象 监听定位消息
	go locate.StartLocate()

	//处理注册函数Handler 用于处理来自/objects/开头地址的请求
	http.HandleFunc("/objects/", objects.Handler)

	//将服务设置在LISTEN_ADDRESS上进行监听
	//os.Getenv("LISTEN_ADDRESS")处理传入参数解析
	log.Fatal(fmt.Sprintf("Err is %v", http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil)))
}

