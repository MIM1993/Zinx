package ziface

/*
	抽象的消息管理模块,存放router集合,使用map数据结构存储msgID和router,实现多路由功能实现
*/

type IMsgHandler interface {
	//添加路由到map集合中
	AddRouter(MsgId uint32, router IRouter)
	//调度路由.根据MsgId
	DoMsgHandler(request IRequest)

	//开启worker工作池
	StartWorkerPool()
	//给worker工作池发数据
	SendMsgToTaskQueue(request IRequest)
}
