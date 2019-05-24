package config

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

//定义全局配置文件的类
type GlobalObj struct {
	Host           string //当前监听的IP
	Port           int    //当前监听的端口号
	Name           string //当前服务器的名字
	Version        string //当前服务器的版本号
	MaxPackageSize int    //读写缓冲区的最大值
}

//定义一个全局的配置文件数据存储容器
var Conf *GlobalObj

//加载配置文件方法
func (g *GlobalObj) LoadConfig() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		fmt.Println("LoadConfig err:", err)
		return
	}

	//将zinx.json中的配置数据解析到conf中，替换掉默认的数据，没替换的依然是默认数据
	err = json.Unmarshal(data, &Conf)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return
	}
}

//导入该模块，就会在执行main函数之前执行init函数
func init() {
	//配置文件读取操作
	//设置默认值
	Conf = &GlobalObj{
		Host:           "0.0.0.0",
		Port:           9999,
		Name:           "Zinx_mim",
		Version:        "ZinxV0.4",
		MaxPackageSize: 512,
	}
	Conf.LoadConfig()
}
