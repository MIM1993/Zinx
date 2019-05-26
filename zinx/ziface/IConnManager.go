package ziface

//定义接口  管理链接  限制用户的增删改查

type IConnManager interface {
	//添加链接
	Add(conn IConnection)
	//删除链接
	Remove(connID uint32)
	//根据链接ID得到链接
	Get(connID uint32) (IConnection, error)
	//得到目前连接的总个数
	Len()uint32
	//清空当前连接
	ClearConn()
}
