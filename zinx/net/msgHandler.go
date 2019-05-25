package net

import (
	"review/zinx/ziface"
	"fmt"
)

//定义路由集合类结构体
type MsgHandler struct {
	//开发者的全部业对应表,消息ID和业务的对应关系
	Apis map[uint32]ziface.IRouter
}

//初始化方法
func NewMsgHandler() ziface.IMsgHandler {
	imsghandler := &MsgHandler{
		Apis: make(map[uint32]ziface.IRouter),
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
