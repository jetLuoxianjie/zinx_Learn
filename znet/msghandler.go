package znet

import (
	"fmt"
	"strconv"
	"zinxLearn/utils"
	"zinxLearn/ziface"
)

type MsgHandle struct {

	// 存放每一个MsgId锁对应的处理方法的map属性
	Apis map[uint32]ziface.IRouter

	// 业务工作worker池的数量
	WorkPoolSize uint32

	// worker负责取任务的消息队列
	TaskQueue []chan ziface.IRequest
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:         make(map[uint32]ziface.IRouter),
		WorkPoolSize: utils.GlobalObject.MaxPacketSize,
		TaskQueue:    make([]chan ziface.IRequest, utils.GlobalObject.MaxPacketSize),
	}
}

// 马上以非阻塞的方式
func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {

	router, ok := mh.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("api msgId = ", request.GetMsgId(), "is not found")
	}

	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)

}

// 为消息添加具体的处理逻辑
func (mh *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	// 1 判断当前的msg绑定的api处理方法是否已经存在
	if _, ok := mh.Apis[msgId]; ok {
		panic("repeated api , msgId = " + strconv.Itoa(int(msgId)))
	}

	// 添加msg与api的绑定关系
	mh.Apis[msgId] = router
	fmt.Println("add api msgIfd = ", msgId)

}

// 启动worker工作池
func (mh *MsgHandle) StartWorkerPool() {
	//遍历需要启动的worker的数量，依次启动
	for i := 0; i < int(mh.WorkPoolSize); i++ {
		// 一个worker被启动
		//给当前的worker对应的任务队列开辟空间
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)

		//启动当前的worker，阻塞的等待对应的任务队列是否有消息传递进来

		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

/**
启动一个worker工作流
*/
func (mh *MsgHandle) StartOneWorker(i int, requests chan ziface.IRequest) {

	fmt.Println("worker id = ", i, " is start..")
	//不断的等待队列中的消息
	for {
		select {
		// 有消息则取出队列的request，并执行绑定的业务方法
		case request := <-requests:
			mh.DoMsgHandler(request)
		}
	}
}

// 将消息交给TaskQueue， 由worker进行处理
func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	// 根据connid来分配当前的连接应该由哪一个worker负责处理
	// 轮训的平均分配法则

	//得到的需要处理的此条连接的workerId
	u := request.GetConnection().GetConnID() % mh.WorkPoolSize
	fmt.Println("add connid :", request.GetConnection().GetConnID(), " request connid")
	mh.TaskQueue[u] <- request

}
