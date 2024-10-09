//Golang实现一个简单的WebServer
package main

import (
	"os"
	"net/http"
)

func main() {
	//声明2个变量
	var dir,port,host string
	//判断参数的长度
	if len(os.Args) == 4 {
		dir = os.Args[1]
		port = os.Args[2]
		host = os.Args[3]
	} else if len(os.Args) == 3 {
		dir = os.Args[1]
		port = os.Args[2]
	} else if len(os.Args) == 2 {
		dir = os.Args[1]
	} else{
		//如果没有参数，则使用默认
		dir = "/etc/ccaa/AriaNg"
		port = "6080"
		host = ""
	}
	
	panic(http.ListenAndServe(host + ":" + port, http.FileServer(http.Dir(dir))))
}