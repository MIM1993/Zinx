package ziface


/*
  抽象的路由模块
 */
type IRouter interface {
	//处理业务之前的方法
	PreHandle(request IRequest)
	//真正处理业务的方法
	Handle(request IRequest)
	//处理业务之后的方法
	PostHandle(request IRequest)
}