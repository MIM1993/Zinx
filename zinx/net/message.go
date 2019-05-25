package net

import "review/zinx/ziface"

/*
	具体的信息包结构体
*/

type Message struct {
	Id      uint32
	Datalen uint32
	Data    []byte
}

//提供一个创建message的函数
func NewMsgPackage(id uint32, data []byte) ziface.IMessage {
	msg := &Message{
		Id:      id,
		Datalen: uint32(len(data)),
		Data:    data,
	}
	return msg
}

//获取函数
func (m *Message) GetMsgLen() uint32 {
	return m.Datalen
}

func (m *Message) GetMsgId() uint32 {
	return m.Id
}

func (m *Message) GetMsgData() []byte {
	return m.Data
}

//setter函数
func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

func (m *Message) SetDatalen(len uint32) {
	m.Datalen = len
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}
