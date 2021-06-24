package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

//分发注册的处理函数
func Handler(w http.ResponseWriter, r *http.Request)  {
	m := r.Method

	if m == http.MethodPut {
		//执行put
		sPut(w, r)
		return
	}

	if m == http.MethodGet {
		//执行get
		sGet(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

//处理Put请求 向本地写入请求的数据内容
func sPut(w http.ResponseWriter, r *http.Request)  {
	//本地存储创建
	fp, err := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer fp.Close()

	//put请求内容写入到创建的本地文件中
	io.Copy(fp, r.Body)
}

//处理Get请求 读取本地文件内容并返回
func sGet(w http.ResponseWriter, r *http.Request) {
	//本地存储读取
	fg, err := os.Open(os.Getenv("STORAGE_ROOT") + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer fg.Close()

	//get内容写入到网页
	io.Copy(w, fg)
}
