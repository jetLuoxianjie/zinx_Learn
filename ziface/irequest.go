package ziface

type IRequest interface {
	GetConnection() IConnection // 获取请求的连接信息
	GetData() []byte            // 获取请求的数据
	GetMsgId() uint32
}
