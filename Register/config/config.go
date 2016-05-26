package config

//配置选项
var RegisterConfig *RegisterConf

const (
	CONFIG_DEFAULT_BOOT_LOG_FILE  = "log/boot.log"
	CONFIG_DEFAULT_ERROR_LOG_FILE = "log/error.log"
)

//Register的配置文件项
type RegisterConf struct {
	Port         uint     //对外提供服务的接口
	OpenHost     string   //开放监听的IP地址
	reciveIPs    []string //接受访问的IP地址配置
	ErrorLogFile string   //普通的错误日志文件
	BootLogFile  string   //启动日志
}

func (rc *RegisterConf) SetReciveIPs() error {
	return nil
}
