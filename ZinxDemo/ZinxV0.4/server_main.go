
package main

import (
	"review/zinx/net"
	"review/zinx/ziface"
	"fmt"
)

//创建路由控制器
type PingRouter struct {
	net.BaseRouter
}

//处理业务之前的方法
func (r *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle ...")
	_,err := request.GetConnection().GetTCPconnection().Write([]byte("before ping ...\n"))
	if err!=nil{
		fmt.Println("call back before ping err:",err)
		return
	}
}

//真正处理业务的方法
func (r *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle ...")
	_,err := request.GetConnection().GetTCPconnection().Write([]byte("ping ping ping ...\n"))
	if err!=nil{
		fmt.Println("call back ping err:",err)
		return
	}
}

//处理业务之后的方法
func (r *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle ...")
	_,err := request.GetConnection().GetTCPconnection().Write([]byte("after ping ...\n"))
	if err!=nil{
		fmt.Println("call back after ping err:",err)
		return
	}
}

func main() {
	//创建一个zinx server对象
	s := net.NewServer("zinx v0.1")

	//添加router
	s.AddRouter(&PingRouter{})
	//让server对象 启动服务
	s.Serve()

	return
}
