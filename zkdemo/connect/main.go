package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func main() {
	// 创建zk连接地址
	// 创建zk连接地址
	hosts := []string{"172.16.1.129:2181"}
	// 连接zk
	conn, _, err := zk.Connect(hosts, time.Second*5)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conn.Server())
}
