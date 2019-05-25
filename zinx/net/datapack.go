package net

import (
	"review/zinx/ziface"
	"bytes"
	"encoding/binary"
)

//具体的拆包对象结构体
type DataPack struct {
}

//创建DataPack对象函数
func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	//数据长度和Id各4个字节,共八字节
	return 8
}

//封包方法   ---- 将 Message  打包成 |datalen|dataID|data|
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	//创建一个存档二进制的字节缓冲区
	dataBuff := bytes.NewBuffer([]byte{})

	//首先,将datalen写进缓冲区中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}

	//然后将dataId写入缓冲区中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	//最后,将Data写入缓冲区中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return nil, err
	}

	//返回这个缓冲区
	return dataBuff.Bytes(), nil
}

//拆包方法  ---  将|datalen|dataID|data|   拆解到 Message 结构体中
func (dp *DataPack) UnPack(binaryData []byte) (ziface.IMessage, error) {
	//定义message结构体,用来进行数据存储-----容器
	msgHead := &Message{}
	//定义一个读取的二进制数据流阅读器
	dataBuff := bytes.NewReader(binaryData)

	//首先读取位二进制流中的datalen							//先msg.Datalen,取到一个整形,然后在取地址
	if err := binary.Read(dataBuff, binary.LittleEndian, &msgHead.Datalen); err != nil {
		return nil, err
	}

	//然后在读取二进制流中的dataId
	if err := binary.Read(dataBuff, binary.LittleEndian, &msgHead.Id); err != nil {
		return nil, err
	}

	//返回值
	return msgHead, nil
}
