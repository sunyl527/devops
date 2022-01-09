package main

import (
	"fmt"
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

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello world</h1>"))

	// 1.接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range r.Header {
		for _, rv := range v {
			w.Header().Set(k, rv)
		}
		fmt.Println(k, v)
	}

	//2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	os.Setenv("VERSION", "test")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)

	//3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	clientIP := getRealIp(r)
	fmt.Println("clientIP:", clientIP)

}

//4.当访问 localhost/healthz 时，应返回 200
func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func getRealIp(r *http.Request) string {
	ip := r.Header.Get("X_REAL_IP")
	if "" == ip {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", homepage)
	mux.HandleFunc("/healthz", healthz)
	if err := http.ListenAndServe("localhost:80", mux); err != nil {
		fmt.Println("error:", err)
	}
}
