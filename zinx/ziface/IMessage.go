package ziface

/*
	将请求的消息封装到message中，定义抽象接口
*/

type IMessage interface {
	//获取消息数据长度
	GetMsgLen() uint32
	//获取消息ＩＤ
	GetMsgId() uint32
	//獲取消息內容
	GetMsgData() []byte

	//設計消息ID
	SetMsgId(uint32)
	//设计数据段长度
	SetDatalen(uint32)
	//设计数据
	SetData([]byte)
}
