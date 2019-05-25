package ziface

type IRequest interface {
	//得到当前请求的连接
	GetConnection() IConnection

	//得到请求的方法
	GetMsg() IMessage
}
