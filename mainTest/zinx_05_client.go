package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinxLearn/znet"
)

func main() {

	fmt.Printf("zinx 05 start..")
	time.Sleep(time.Second)

	dial, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("dail error ", err)
		return
	}

	for {
		pack := znet.NewDataPack()

		msgPackage := znet.NewMsgPackage(12, []byte("hello world!"))
		bytes, err := pack.Pack(msgPackage)
		if err != nil {
			fmt.Println("封包错误 。。。")
			return
		}
		_, err = dial.Write(bytes)
		if err != nil {
			fmt.Println("ping write error ")
			return
		}
		// 先读取流中的数据head部分
		headData := make([]byte, pack.GetHeadLen())
		_, err = io.ReadFull(dial, headData)
		if err != nil {
			fmt.Println("write data error")
			return
		}
		msg, err := pack.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error ")
			return
		}
		if msg.GetDataLen() > 0 {
			// msg是有data数据的需要再次读取数据
			message := msg.(*znet.Message)
			message.Data = make([]byte, msg.GetDataLen())

			_, err := io.ReadFull(dial, message.Data)
			if err != nil {
				fmt.Println("server unpack data error ", err)
				return
			}

			fmt.Println("==> Recv Msg: ID=", message.Id, ", len=", message.DataLen, ", data=", string(message.Data))

		}
		time.Sleep(1 * time.Second)
	}
}
