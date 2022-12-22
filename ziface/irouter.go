package ziface

/**
路由接口，这里面路由是使用框架者给连接自定义的处理业务的方法
路由里的IRequest则包含该链接的连接消息和改连接的请求数据信息
*/
type IRouter interface {
	PreHandle(request IRequest)  // 在处理conn业务之前的狗子方法
	Handle(request IRequest)     // 处理conn业务的方法
	PostHandle(request IRequest) // 处理conn业务之后的钩子方法

}
