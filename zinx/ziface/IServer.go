package ziface

type IServer interface {
	//启动服务器
	Start()
	//停止服务器
	Stop()
	//运行服务
	Serve()
	//添加路由函数
	AddRouter(MsgId uint32, router IRouter)

	//提供一个方法,获取链接管理模块
	GetConnMgr() IConnManager

	//注册函数--->创建链接之后需要调用的HOOK方法
	AddOnConnStart(hookFunc func(conn IConnection))
	//注册函数--->销毁链接之前需要调用的HOOK方法
	AddOnConnStop(hookFunc func(conn IConnection))
	//调用创建链接后运行的方法
	CallOnConnStart(conn IConnection)
	//调用销毁链接前运行的方法
	CallOnConnStop(conn IConnection)
}
