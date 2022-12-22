package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"zinxLearn/ziface"
)

/**
存储一些有段zinx框架的全局的参数，供其他模块使用
一些参数也可以通过用户根据zinx.json来配置
*/

type GlobalObj struct {
	TcpServet ziface.IServer // 当前的全局server对象

	Host string // 当前服务器主机ip

	TcpPort int // 当前的服务器监听的端口

	Name string // 当前服务器的名称

	Version string // 当前的版本号

	MaxPacketSize uint32 // 都需数据包的最大值
	MaxConn       int    // 当前服务器主机允许的最大连接个数

	WorkerPoolSize   uint32 //业务工作Worker池的数量
	MaxWorkerTaskLen uint32 //业务工作Worker对应负责的任务队列最大任务存储数量

	// config file path
	ConfFilePath string
}

/**
定义一个全局的GlobalObject对象

*/

var GlobalObject *GlobalObj

// 读取用户的配置表属性
func (g *GlobalObj) Reload() {
	file, err := ioutil.ReadFile(g.ConfFilePath)
	if err != nil {
		panic(err)
	}

	// 将json数据解析到struct中
	fmt.Println("json data : ", string(file))
	err = json.Unmarshal(file, &GlobalObject)
	fmt.Println("data globalObject : ", &GlobalObject)
	if err != nil {
		panic(err)
	}

}

// 提供init方法

func init() {
	// 初始化GlobalObject变量，设置一些默认值
	GlobalObject = &GlobalObj{
		Name:             "zinxServerApp",
		Version:          "v0.4",
		TcpPort:          8081,
		Host:             "0.0.0.0",
		MaxConn:          12000,
		MaxPacketSize:    512,
		ConfFilePath:     "conf/zinx.json",
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
	}

	// 从配置表中加载用户的配置
	GlobalObject.Reload()

}
