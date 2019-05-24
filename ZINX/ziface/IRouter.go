package ziface

//定义路由接口  抽象的路由
type IRouter interface {
	//处理业务之前的方法
	PreHandle(request IRequest)
	//处理真正的处理方发
	Handle(request IRequest)
	//处理业务之后的方法
	PostHandle(request IRequest)
}
