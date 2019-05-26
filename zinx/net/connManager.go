package net

import (
	"review/zinx/ziface"
	"sync"
	"fmt"
	"errors"
)

//定义链接管理结构体
type ConnManager struct {
	//管理全部的链接
	connections map[uint32]ziface.IConnection
	//锁
	connLock sync.RWMutex
}

//创建链接管理结构体方法
func NewConnManager() ziface.IConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

//添加链接
func (connMgr *ConnManager) Add(conn ziface.IConnection) {
	//加锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	connMgr.connections[conn.GetConnID()] = conn

	fmt.Println("Add connid = ", conn.GetConnID(), "to manager succ!!")
}

//删除链接
func (connMgr *ConnManager) Remove(connID uint32) {
	//加锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	delete(connMgr.connections, connID)
	fmt.Println("Remove connid = ", connID, " from manager succ!!")
}

func (connMgr *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	//加读锁
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()

	if conn, ok := connMgr.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not FOUND!")
	}
}

//获取目前连接的总个数
func (connMgr *ConnManager) Len() uint32 {
	return uint32(len(connMgr.connections))
}

//清空当前链接
func (connMgr *ConnManager) ClearConn() {
	//加锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//遍历全部connection
	for connID, conn := range connMgr.connections {
		//将全部的conn链接关闭
		conn.Stop()
		//删除链接
		delete(connMgr.connections, connID)
	}

	fmt.Println("Clear All Conections succ! conn num = ", connMgr.Len())
}
