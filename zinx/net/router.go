package net

import "review/zinx/ziface"

/*
	具体的路由
*/

type BaseRouter struct {
}

//处理业务之前的方法
func (r *BaseRouter) PreHandle(request ziface.IRequest) {

}

//真正处理业务的方法
func (r *BaseRouter) Handle(request ziface.IRequest) {

}

//处理业务之后的方法
func (r *BaseRouter) PostHandle(request ziface.IRequest) {

}
