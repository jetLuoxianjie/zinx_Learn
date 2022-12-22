package main

import (
	"fmt"
	"io"
	"net"
	"zinxLearn/znet"
)

// 只是负责测试封包和拆包的测试
func main() {
	// 创建socket tcp server
	listen, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil {
		fmt.Println("listen error :", err)
	}
	// 创建服务器gogroutine,负责从客户端读取粘包数据负责解析

	fmt.Println(listen)

	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("listen error : ", err)
			return
		}
		go func(conn net.Conn) {
			// 创建封包拆包的对象
			dp := znet.NewDataPack()
			for {
				// 读取head中的8个字节长度
				bytes := make([]byte, dp.GetHeadLen())
				// ReadFull 会把msg填充满为止
				_, err2 := io.ReadFull(conn, bytes)
				if err2 != nil {
					fmt.Println("read full error : ", err2)
					break
				}

				unpack, err2 := dp.Unpack(bytes)
				if err2 != nil {
					fmt.Println("server unpack err: ", err2)
					return
				}

				if unpack.GetDataLen() > 0 {
					message := unpack.(*znet.Message)
					message.Data = make([]byte, message.GetDataLen())

					_, err2 := io.ReadFull(conn, message.Data)
					if err2 != nil {
						fmt.Println("uppack data error ", err2)
						return
					}
					fmt.Println("==> Recv Msg: ID=", message.Id, ", len=", message.DataLen, ", data=", string(message.Data))
				}

			}
		}(accept)
	}
}
