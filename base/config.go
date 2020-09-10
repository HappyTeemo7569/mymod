package base

import (
	"gopkg.in/gcfg.v1"
)

func InitConfig() {
	initBaseConfig()
}

var configBase tagConfigBase

type tagConfigBase struct {
	MySql tagMysqlConfig
	Redis tagRedisConfig
	Log   tagLogConfig
}

type tagMysqlConfig struct {
	Auth string
	Pwd  string
	Addr string
	Port int
	Db   string
}

type tagRedisConfig struct {
	Addr string
	Port int
	Pwd  string
	Rpc  int
	Api  int
}

type tagLogConfig struct {
	Path string
}

//初始化基础配置
func initBaseConfig() {
	err := gcfg.ReadFileInto(&configBase, "config/configBase.ini")
	if err != nil {
		panic("模块加载错误：配置文件读取错误：" + err.Error())
		return
	}

}
