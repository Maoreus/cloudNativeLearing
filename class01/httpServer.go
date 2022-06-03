package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

/*
接收客户端 request，并将 request 中带的 header 写入 response header
读取当前系统的环境变量中的 VERSION 配置，并写入 response header
Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
当访问 localhost/healthz 时，应返回 200
*/

type MyResponse struct {
	code    int
	data    map[string]string
	message string
}

func healthzHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Printf("start handle request......")
	//获取version写入header
	headerCopy := req.Header.Clone()
	version := os.Getenv("VERSION")
	headerCopy.Add("version", version)

	ip := fetchClientIP(req)
	data := make(map[string]string, 2)
	data["ip"] = ip
	ret := MyResponse{code: http.StatusOK, data: data, message: "success"}
	jsonOut, err := json.Marshal(ret)
	writer.WriteHeader(http.StatusOK)
	if err != nil {
		writer.Write(jsonOut)

	}
	fmt.Print("finish......")

}

// 获取clientIp
func fetchClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

func main() {
	//注册路由
	http.HandleFunc("/healthz", healthzHandler)
	//开启监听8080
	listen := http.ListenAndServe(":8080", nil)
	if listen != nil {
		log.Fatal("listen to 8080 port error！ {}", listen)

	}
}
