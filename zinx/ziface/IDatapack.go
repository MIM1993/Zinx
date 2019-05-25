package ziface

/*
	封包拆包接口
*/

type IDatapack interface {
	//获取二进制文件的包头  固定长度返回8
	GetHeadLen() uint32
	//封包方法   将Message 打包成 |datalen|dataId|data|
	Pack(msg IMessage) ([]byte, error)
	//拆包方法  将|datalen|dataId|data| 拆解到Message 结构体中
	UnPack([]byte) (IMessage, error)
}
