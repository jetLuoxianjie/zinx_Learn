package main

import (
	"fmt"
	"net"
	"zinxLearn/znet"
)

//只负责发送粘包数据的测试客户端
func main() {

	// 客户端gortine，负责模拟粘包的数据
	dial, err := net.Dial("tcp", "127.0.0.1:8082")
	if err != nil {
		fmt.Println("net dial error : ", err)
	}

	pack := znet.NewDataPack()

	msg1 := &znet.Message{
		Id:      0,
		DataLen: 5,
		Data:    []byte("hello"),
	}
	msg2 := &znet.Message{
		Id:      1,
		DataLen: 7,
		Data:    []byte("worldop"),
	}

	// 压缩数据msg1和msg2
	sendData1, err := pack.Pack(msg1)
	if err != nil {
		fmt.Println("msg1 pack error: ", err)
	}
	sendData2, err := pack.Pack(msg2)
	if err != nil {
		fmt.Println("msg2 pack error: ", err)
	}

	//将数据进行打包
	bytes := append(sendData1, sendData2...)

	dial.Write(bytes)
	select {}
}
