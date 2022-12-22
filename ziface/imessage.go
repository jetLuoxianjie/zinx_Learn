package ziface

/**
将请求的一个消息封装到message中，定义抽象层接口
*/

type IMessage interface {
	GetDataLen() uint32 // 获取数据的长度
	GetMsgId() uint32   // 获取消息的id
	GetData() []byte    // 获取消息内容

	SetMsgId(uint322 uint32)   // 设置消息的id
	SetData([]byte)            // 设置消息的内容
	SetDataLen(uint322 uint32) // 设置消息的数据长度

}
