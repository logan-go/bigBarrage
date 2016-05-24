package main

//配置选项
var RegisterConfig *RegisterConf

//Register的配置文件项
type RegisterConf struct {
	Port      uint     //对外提供服务的接口
	OpenHost  string   //开放监听的IP地址
	reciveIPs []string //接受访问的IP地址配置
}
