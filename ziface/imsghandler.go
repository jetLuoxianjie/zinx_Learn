package ziface

/**
消息管理的抽象层
这里面有两个方法，AddRouter()就是添加一个msgId和一个路由关系到Apis中，
那么DoMsgHandler()则是调用Router中具体Handle()等方法的接口。

*/

type IMsgHandle interface {

	// 马上以非阻塞的方式处理消息
	DoMsgHandler(request IRequest)

	//为消息添加具体的处理逻辑
	AddRouter(msgId uint32, router IRouter)

	// 启动worker工作池
	StartWorkerPool()

	// 将消息交给taskQueue，有worker进行处理
	SendMsgToTaskQueue(reader IRequest)
}
