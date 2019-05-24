package ziface

//封装请求接口
type IRequest interface {
	//得到当前请求的链接
	GetConnection() IConnection
	//得到链接的数据
	GetData() []byte
	//得到链接的长度
	GetLen() int
}
