package main

import (
	"zinx/ZINX/znet"
	"zinx/ZINX/ziface"
	"fmt"
)

type PingRouter struct {
	znet.BaseRouter
}



func(p *PingRouter)Handle(request ziface.IRequest){
	fmt.Println("server call back is ...")
	_,err :=request.GetConnection().GetTCPConnection().Write([]byte("hello client"))
	if err!=nil{
		fmt.Println("Handle call back err:",err)
		return
	}
}

func main(){
	s:=znet.NewServer("zinx_v0.1")

	//添加路由起器
	s.AddRouter(&PingRouter{})

	s.Serve()
}
