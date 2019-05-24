package znet

import "zinx/ZINX/ziface"

//定义router类型结构体   具体的路由
type BaseRouter struct {
}

//业务处理之前操作函数
func (r *BaseRouter) PreHandle(request ziface.IRequest) {
}

//处理业务函数
func (r *BaseRouter) Handle(request ziface.IRequest) {
}

//处理业务之后操作的函数
func (r *BaseRouter) PostHandle(request ziface.IRequest) {
}
