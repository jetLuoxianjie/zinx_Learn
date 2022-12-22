package main

import (
	"fmt"
	"zinxLearn/ziface"
	"zinxLearn/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

//Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgId, ", data=", string(request.GetData()))

	//回写数据
	err := request.GetConnection().SendMsg(234, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

//HelloZinxRouter Handle
type HelloZinxRouter struct {
	znet.BaseRouter
}

func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgId(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(112, []byte("Hello Zinx Router V0.6"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	// 创建服务器的句柄
	server := znet.NewServer("zinx_05")

	// 添加路由
	router := PingRouter{}
	server.AddRouter(12, &router)

	zinxRouter := HelloZinxRouter{}
	server.AddRouter(13, &zinxRouter)

	server.Server()
}
