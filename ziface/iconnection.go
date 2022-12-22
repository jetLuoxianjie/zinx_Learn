package ziface

import "net"

/**
V0.1版本我们已经实现了一个基础的Server框架，现在我们需要对客户端链接和不同的客户端链接所处理的不同业务再做一层接口封装，当然我们先是把架构搭建起来。
*/

type IConnection interface {

	// 启动连接， 让当前连接开始工作
	Start()

	// 停止连接，结束当前的连接状态M
	Stop()

	// 从当前连接获取原始的socket tcpConn
	GetTCPConnection() *net.TCPConn

	// 获取档期啊能连接诶的ID
	GetConnID() uint32

	// 获取远程客户端地址信息
	RemoteAddr() net.Addr

	// 5.0直接将Message数据发送给远程的tcp客户端
	SendMsg(msgId uint32, data []byte) error
}

/**
该接口的一些基础方法，代码注释已经介绍的很清楚，这里先简单说明一个HandFunc这个函数类型，
这个是所有conn链接在处理业务的函数接口，第一参数是socket原生链接，
第二个参数是客户端请求的数据，第三个参数是客户端请求的数据长度。这样，如果我们想要指定一个conn的处理业务，
只要定义一个HandFunc类型的函数，然后和该链接绑定就可以了。
*/
// 定义一个统一处理连接业务的接口
type HandFunc func(*net.TCPConn, []byte, int) error
