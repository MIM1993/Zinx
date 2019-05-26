package net

import (
	"review/zinx/ziface"
	"fmt"
	"review/zinx/config"
)

//定义路由集合类结构体
type MsgHandler struct {
	//开发者的全部业对应表,消息ID和业务的对应关系
	Apis map[uint32]ziface.IRouter
	//业务工作池的数量
	WorkerPoolSize uint32
	//消息队列  切片 IRequest
	TaskQueue []chan ziface.IRequest
}

//初始化方法
func NewMsgHandler() ziface.IMsgHandler {
	imsghandler := &MsgHandler{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: config.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, config.GlobalObject.WorkerPoolSize),
	}
	return imsghandler
}

//添加路由到map映射表中
func (mh *MsgHandler) AddRouter(MsgId uint32, router ziface.IRouter) {
	//判断新添加的key是否重复
	if _, ok := mh.Apis[MsgId]; ok {
		//msgId已经注册
		fmt.Println("repeat Api MsgId=", MsgId)
		return
	}
	//添加msgId和router
	mh.Apis[MsgId] = router
	fmt.Println("Api MsgId =", MsgId, "succ!")
}

//调用函数  根据MsgId
func (mh *MsgHandler) DoMsgHandler(request ziface.IRequest) {
	//根据Request获取MsgId
	MsgId := request.GetMsg().GetMsgId()

	//获取对应router
	Router, ok := mh.Apis[MsgId]
	if !ok {
		fmt.Println("Api MsgId=", MsgId, "Not Found! Need Add!")
		return
	}

	//根据MsgId,找到对应的router进行调用
	Router.PreHandle(request)
	Router.Handle(request)
	Router.PostHandle(request)
}

//开启worker工作池
func (mh *MsgHandler) StartWorkerPool() {
	fmt.Println("Worker Pool is Started...")

	//根据WorkerSize来创建worker  groutine
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		//开启一个worker groutine

		//给当前的worker groutine 所绑定的channel对象开辟空间  0号worker-----0号channel
		mh.TaskQueue[i] = make(chan ziface.IRequest, config.GlobalObject.MaxWorkerTaskLen)

		//启动当前worker
		go mh.startOneWorker(i, mh.TaskQueue[i])
	}
}

//真正处理worker业务的groutine
func (mh *MsgHandler) startOneWorker(workerID int, taskQueue chan ziface.IRequest) {
	fmt.Println("worker ID =", workerID, "is starting...")

	//等待消息传递过来
	for {
		select {
		case req := <-taskQueue:
			mh.DoMsgHandler(req)
		}
	}
}

//给worker工作池发数据
func (mh *MsgHandler) SendMsgToTaskQueue(request ziface.IRequest) {
	//1 将消息 平均分配给worker 确定当前的request到底要给哪个worker来处理
	//1个客户端绑定一个worker来处理
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	//2 直接将 request 发送给对应的worker的taskqueue
	mh.TaskQueue[workerID] <- request
}
