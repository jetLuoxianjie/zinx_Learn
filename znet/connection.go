package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinxLearn/ziface"
)

// 实现 iconnection接口
type Connection struct {

	// 当前的socket tcp连接
	Conn *net.TCPConn

	//当前的ID 也可以成为SessionID 全局唯一
	ConnID uint32

	// 当前连接的关闭状态
	isClosed bool

	////2.0版本替换 处理连接的api方法
	//handleAPI ziface.HandFunc

	//Router ziface.IRouter
	MsgHandle ziface.IMsgHandle

	// 告知该连接已经退出/停止的channel
	ExitBuffChan chan bool

	// 版本7.0 无缓冲的管道，用于读写两个groutine之间的消息通信
	msgChan chan []byte
}

// 创建连接的方法
func NewConntion(conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandle) *Connection {

	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		MsgHandle:    msgHandler,
		ExitBuffChan: make(chan bool, 1),
		msgChan:      make(chan []byte), // msgChan初始化
	}
	return c
}

/**
写消息Goroutine, 用户数据发送给客户端
*/
func (c *Connection) StartWriter() {
	fmt.Println("[Writer Groutine is running]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn writer exit]")

	for {
		select {
		case data := <-c.msgChan:
			// 有数据要写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data error :, ", err, " Conn writer exit")
				return
			}
		case <-c.ExitBuffChan:
			// conn已经关闭
			return
		}
	}
}

//处理当前数据的Groutine
//noinspection GoInvalidCompositeLiteral
func (c *Connection) StartReader() {
	fmt.Println(" 开始执行 reader Groutine 方法 。。。。")
	defer fmt.Println(c.Conn.RemoteAddr().String(), " 连接 reader exit!")

	defer c.Stop()

	for {

		// 版本5.0 创建封包拆包对象
		dp := NewDataPack()

		//读取我们最大的数据
		buf := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), buf)
		if err != nil {
			fmt.Println("read data error ...")
			c.ExitBuffChan <- true
			continue
		}

		// 调用当前连接业务（执行当前的conn绑定的handle方法）
		// 2.0 版本err = c.handleAPI(c.Conn, buf, read)

		// 版本5.0 将获取的数据进行拆包
		//拆包，得到msgid 和 datalen 放在msg中
		msg, err := dp.Unpack(buf)
		if err != nil {
			fmt.Println("unpack error ", err)
			c.ExitBuffChan <- true
			continue
		}

		//根据 dataLen 读取 data，放在msg.Data中
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error ", err)
				c.ExitBuffChan <- true
				continue
			}
		}
		msg.SetData(data)

		//3.0版本IRouter
		var request = Request{
			conn: c,
			msg:  msg,
		} // 执行注册的路由方法
		fmt.Println("request data : ", request.GetConnection().GetTCPConnection())
		fmt.Println(request.GetConnection().GetConnID())
		req := Request{
			conn: c,
			msg:  msg,
		}
		//从路由Routers 中找到注册绑定Conn的对应Handle
		//go func(request ziface.IRequest) {
		//	//执行注册的路由方法
		//	c.Router.PreHandle(request)
		//	c.Router.Handle(request)
		//	c.Router.PostHandle(request)
		//}(&req)

		// 版本 6.0多路由模式
		go c.MsgHandle.DoMsgHandler(&req)

	}

}

// 新增5.0Message数据发送给远程的tcp客户端
/**
版本7.0 Reader讲发送客户端的数据改为发送至Channel
*/
func (r *Connection) SendMsg(msgId uint32, data []byte) error {
	if r.isClosed {
		return errors.New("conntion closed when send message")
	}

	// 将data数据封包
	dp := NewDataPack()
	pack, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("PACK error msg id = ", msgId)
		return errors.New("pack error msg")
	}

	// 写回客户端
	//_, err = r.Conn.Write(pack)
	//if err != nil {
	//	r.ExitBuffChan <- true
	//	return errors.New("conn write error ")
	//}

	// 版本7.0直接写到消息通道中，不直接写会客户端
	r.msgChan <- pack

	return nil

}

// 启动连接，让连接开始工作
func (c *Connection) Start() {

	// 开始处理该链接读取客户端数据之后的请求业务
	go c.StartReader()

	// 版本7.0开始写数据到goroutine中
	go c.StartWriter()
	for {
		select {
		case <-c.ExitBuffChan:
			fmt.Println("============exit=========")
			// 得到消息退出，不用阻塞
			return
		}
	}

}

// 停止连接，结束当前连接
func (c *Connection) Stop() {
	// 1. 如果当前连接已经关闭
	if c.isClosed == true {
		return
	}

	c.isClosed = true
	//TODO Connection Stop() 如果用户注册了该链接的关闭回调业务，那么在此刻应该显示调用

	c.Conn.Close()

	// 通知从缓冲队列读取数据的业务， 该链接已经关闭
	c.ExitBuffChan <- true

	// 关闭连接全部通道
	close(c.ExitBuffChan)

}

// 从当前连接获取原始连接的socket TCPConn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 从当前的获取连接的ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID

}

// 获取远程的客户端的地址
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
