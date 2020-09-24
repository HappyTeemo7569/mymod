package base

func init() {

	InitConfig()

	//连接池
	InitLog()
	Logger.Infof("配置初始化--日志初始化成功")

	InitRedis()
	Logger.Infof("配置初始化--redis初始化成功")

	InitMysqlNormal()
	InitMysqlNormalGOrm()
	Logger.Infof("配置初始化--mysql初始化成功")

}
