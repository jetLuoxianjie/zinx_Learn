package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinxLearn/utils"
	"zinxLearn/ziface"
)

/**
实现拆包封包类
*/

// 封包拆包的实例，暂时不需要成员
type DataPack struct {
}

// 封包拆包的实例方法初始化
func NewDataPack() *DataPack {
	return &DataPack{}
}

// 获取包头的长度方法
func (dp *DataPack) GetHeadLen() uint32 {
	// id uint32(4字节) + datalen uint32(4字节)
	//todo 这里读取字节的大小是写死的
	return 8
}

// 封包方法（压缩数据）
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// 创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	// 写dataLen
	err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen())
	if err != nil {
		return nil, err
	}

	// 写msgId
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId())
	if err != nil {
		return nil, err
	}

	// 写data数据
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetData())
	if err != nil {
		return nil, err
	}

	//返回数据
	return dataBuff.Bytes(), nil
}

// 拆包方法（解压数据）
func (dp *DataPack) Unpack(data []byte) (ziface.IMessage, error) {
	// 创建一个从输入二进制数据的ioReader
	reader := bytes.NewReader(data)

	// 只解压head信息，得到datalen和msgid
	m := &Message{}

	// 读取datalen 固定读取8个字节，todo 因为uint32占4个字节
	err := binary.Read(reader, binary.LittleEndian, &m.DataLen)
	if err != nil {
		return nil, err
	}

	// 读取msgid
	err = binary.Read(reader, binary.LittleEndian, &m.Id)
	if err != nil {
		return nil, err
	}

	// 读取data，判断data的长度是否超出我们的最大包的长度
	if utils.GlobalObject.MaxPacketSize > 0 && m.DataLen > utils.GlobalObject.MaxPacketSize {
		return nil, errors.New("Too large msg data recieved")
	}

	// 这里我们只需要吧head的数据拆包就可以了，然后通过head的长度，再从conn读取一次数据
	return m, nil
}
