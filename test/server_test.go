package test

import (
	"fmt"
	"net"
	"testing"
	"time"
	"zinxLearn/ziface"
	"zinxLearn/znet"
)

/**
模拟客户端
*/

func ClientTest() {

	fmt.Println("Client Test ... start")
	// 3秒之后发起测试请求
	time.Sleep(3 * time.Second)

	dial, err := net.Dial("tcp", "127.0.0.1:8081")

	if err != nil {
		fmt.Println("clint start err, exit!")
		return
	}

	for {
		_, err := dial.Write([]byte("hello zinx"))
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		bytes := make([]byte, 512)

		read, err := dial.Read(bytes)
		if err != nil {
			fmt.Println(" read buf error: ", bytes, " ", read)
			return
		}

		fmt.Println(" server call back ", string(bytes[:read]), read)
		time.Sleep(time.Second)

	}

}

type MyRouter struct {
	znet.BaseRouter
}

func (router *MyRouter) PreHandle(req ziface.IRequest) {
	fmt.Println(" call router Prehandle..")
	_, err := req.GetConnection().GetTCPConnection().Write([]byte("preHandle run ...." +
		time.Now().Format("2006-01-02 15:04:05")))
	if err != nil {
		fmt.Println(" preHandle error")
	}

}

func (router *MyRouter) Handle(req ziface.IRequest) {

}

func (router *MyRouter) PostHandle(req ziface.IRequest) {

}

func TestServer(t *testing.T) {
	/**
	服务器测试
	*/

	server := znet.NewServer("myZinx")

	router := MyRouter{}
	server.AddRouter(&router)

	go ClientTest()

	server.Server()

}
