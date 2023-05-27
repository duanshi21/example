package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

var (
	hosts       = []string{"127.0.0.1:2181"}
	path        = "/gowork2"
	flags int32 = zk.FlagEphemeral
	data        = []byte("zk usedata2")
	acls        = zk.WorldACL(zk.PermAll)
)

func main() {
	// 连接zk
	conn, _, err := zk.Connect(hosts, time.Second*5)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 开始监听path
	_, _, event, err := conn.ExistsW(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 协程调用监听事件
	go watchZkEvent(event)

	// 触发创建数据操作
	create(conn, path, data)

}

// zk 回调函数
func watchZkEvent(e <-chan zk.Event) {
	event := <-e
	fmt.Println("###########################")
	fmt.Println("path: ", event.Path)
	fmt.Println("type: ", event.Type.String())
	fmt.Println("state: ", event.State.String())
	fmt.Println("---------------------------")
}

// 创建数据
func create(conn *zk.Conn, path string, data []byte) {
	_, err := conn.Create(path, data, flags, acls)
	if err != nil {
		fmt.Printf("创建数据失败: %v\n", err)
		return
	}
	fmt.Println("创建数据成功")
}
