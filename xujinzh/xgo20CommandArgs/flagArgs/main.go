package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义几个变量，用于接收命令行参数
	var user string
	var passwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&passwd, "p", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")

	// 必须调用转换
	flag.Parse()

	// 输出结果
	fmt.Printf("user=%v\npassword=%v\nhost=%v\nport=%v\n", user, passwd, host, port)
}
