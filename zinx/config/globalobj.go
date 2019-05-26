package config

import (
	"io/ioutil"
	"encoding/json"
)

//全局配置文件的类
type GlobalObj struct {
	//server

	Host string //当前监听的IP
	Port int    //当前监听的端口号
	Name string //当前服务器的名字

	Version        string //版本号
	MaxPackageSize int    //每一次read一次的最大长度

	WorkerPoolSize   uint32 //当前服务器开启了多少worker Groutine
	MaxWorkerTaskLen uint32 //每个worker的对应的消息对象  容许缓存的最大任务数量

	MaxConn uint32 //当前server的最大链接数量
}

//定义一个全局的对外的配置文件的对象
var GlobalObject *GlobalObj

//添加一个加载配置文件的方法
func (g *GlobalObj) LoadConfig() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}

	//将zinx.json的数据转化的到 GlobalObject中， json一个解析过程
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

//只要import 当前模块 就会执行init方法 加载配置文件

func init() {
	//配置文件的读取操作
	GlobalObject = &GlobalObj{
		//设置默认值
		Name:             "ZinxServerApp",
		Host:             "0.0.0.0",
		Port:             8999,
		Version:          "V0.4",
		MaxPackageSize:   512,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 4096,
		MaxConn:1000,
	}

	//加载文件函调用
	GlobalObject.LoadConfig()
}
