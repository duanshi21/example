package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

var (
	path = "/go"
)

func main() {
	// 创建zk连接地址
	hosts := []string{"127.0.0.1:2181"}
	// 连接zk
	conn, _, err := zk.Connect(hosts, time.Second*5)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conn.Server())

	// 增删改查
	//add(conn)
	get(conn)
	modify(conn)
	//del(conn)
	get(conn)
}

// 增加
func add(conn *zk.Conn) {
	var data = []byte("golang value")
	// flags有四种值：
	// 0：永久，除非手动删除
	// zk.FlagEphemeral = 1:短暂，session断开则该节点也被删除
	// zk.FlagSequence = 2:会自动在节点后面添加序号
	// 3:Ephemeral和Sequence，即，短暂且自动添加序号
	var flags int32 = 0
	// 获取访问控制权限
	acls := zk.WorldACL(zk.PermAll)
	s, err := conn.Create(path, data, flags, acls)
	if err != nil {
		fmt.Printf("创建失败：%v\n", err)
		return
	}
	fmt.Printf("创建：%s 成功", s)
}

// 查看
func get(conn *zk.Conn) {
	data, _, err := conn.Get(path)
	if err != nil {
		fmt.Printf("查询%s失败，err：%v\n", path, err)
		return
	}
	fmt.Printf("%s 的值为 %s\n", path, string(data))
}

// 修改与增加不同在于其函数中的version参数，其中version是用于CAS支持
// 可以通过此种方式保证原子性
// 修改
func modify(conn *zk.Conn) {
	newData := []byte("hello go-zookeeper")
	_, sate, _ := conn.Get(path)
	_, err := conn.Set(path, newData, sate.Version)
	if err != nil {
		fmt.Printf("数据修改失败：%v\n", err)
		return
	}
	fmt.Println("数据修改成功")
}

// 删除
func del(conn *zk.Conn) {
	_, sate, _ := conn.Get(path)
	err := conn.Delete(path, sate.Version)
	if err != nil {
		fmt.Printf("数据删除失败：%v\n", err)
		return
	}
	fmt.Println("数据删除成功")
}
