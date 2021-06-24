package main

import (
	"fmt"
	"log"
	"myStorage/objects"
	"net/http"
	"os"
)

func main()  {
	//处理注册函数Handler 用于处理来自/objects/开头地址的请求
	http.HandleFunc("/objects/", objects.Handler)

	//将服务设置在LISTEN_ADDRESS上进行监听
	//os.Getenv("LISTEN_ADDRESS")处理传入参数解析
	log.Fatal(fmt.Sprintf("Err is %v", http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil)))
}

