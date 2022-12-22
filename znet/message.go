package znet

type Message struct {
	Id uint32 // 消息的id

	DataLen uint32 // 消息的长度

	Data []byte // 消息的内容

}

// 创建一个Message消息包
func NewMsgPackage(id uint32, data []byte) *Message {

	return &Message{
		id,
		uint32(len(data)),
		data,
	}
}

// 获取消息数据段长度
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

// 获取消息的id
func (msg *Message) GetMsgId() uint32 {
	return msg.Id
}

// 获取消息的内容
func (msg *Message) GetData() []byte {
	return msg.Data
}

// 设置消息
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}

// 设置消息的id
func (msg *Message) SetMsgId(id uint32) {
	msg.Id = id
}

// 设置消息内容
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
