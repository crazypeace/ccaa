//Golang实现一个简单的WebServer
package main

import (
  "flag"
  "strconv"
  "net/http"
)

var (
  dir string
  host string
  port int
)

func init() {
  // 定义命令行参数（参数名，默认值，参数说明）
  flag.StringVar(&dir, "d", "/etc/ccaa/AriaNg", "Web directory")
  flag.StringVar(&host, "h", "", "Listening host (default listen *)")
  flag.IntVar(&port, "p", 6080, "Listening port")
}

func main() {
  // 解析命令行参数
  flag.Parse()

  panic(http.ListenAndServe(host + ":" + strconv.Itoa(port), http.FileServer(http.Dir(dir))))
}