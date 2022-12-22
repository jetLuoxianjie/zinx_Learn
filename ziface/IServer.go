package ziface

type IServer interface {

	//启动
	Start()

	//停止
	Stop()

	//服务开启
	Server()

	// 版本3.0 路由功能： 给当前服务注册一个路由业务方法，供客户端连接处理使用
	// 6.0添加多路由
	AddRouter(msgId uint32, router IRouter)
}
