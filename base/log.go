package base

import (
	go_logger "github.com/phachon/go-logger"
)

//日志
var Logger *go_logger.Logger

func InitLog() {

	Logger = go_logger.NewLogger()

	Logger.Detach("console")

	// 命令行输出配置
	consoleConfig := &go_logger.ConsoleConfig{
		Color:      true,                                                                         // 命令行输出字符串是否显示颜色
		JsonFormat: false,                                                                        // 命令行输出字符串是否格式化
		Format:     "[%timestamp_format%] [%level_string%] [%file% %function% %line%] [%body%] ", // 如果输出的不是 json 字符串，JsonFormat: false, 自定义输出的格式
	}
	//"[%timestamp_format%] [%level_string%] [%file% %function% %line%] %body% "

	// 添加 console 为 Logger 的一个输出
	Logger.Attach("console", go_logger.LOGGER_LEVEL_DEBUG, consoleConfig)

	// 文件输出配置
	fileConfig := &go_logger.FileConfig{
		Filename: configBase.Log.Path + "/test.log", // 日志输出文件名，不自动存在
		// 如果要将单独的日志分离为文件，请配置LealFrimeNem参数。
		LevelFileName: map[int]string{
			Logger.LoggerLevel("Error"):     configBase.Log.Path + "/Error.log", // Errorf 级别日志被写入 error .log 文件
			Logger.LoggerLevel("Info"):      configBase.Log.Path + "/Info.log",  // Infof 级别日志被写入到 info.log 文件中
			Logger.LoggerLevel("Debug"):     configBase.Log.Path + "/Debug.log", // Debugf 级别日志被写入到 debug.log 文件中
			Logger.LoggerLevel("Emergency"): configBase.Log.Path + "/Emergency.log",
			Logger.LoggerLevel("Alert"):     configBase.Log.Path + "/Alert.log",
			Logger.LoggerLevel("Critical"):  configBase.Log.Path + "/Critical.log",
			Logger.LoggerLevel("Notice"):    configBase.Log.Path + "/Notice.log",
			Logger.LoggerLevel("Warning"):   configBase.Log.Path + "/Warning.log",
		},
		MaxSize:    1024 * 1024 * 10,                                                            // 文件最大值（KB），默认值0不限
		MaxLine:    100000,                                                                      // 文件最大行数，默认 0 不限制
		DateSlice:  "d",                                                                         // 文件根据日期切分， 支持 "Y" (年), "m" (月), "d" (日), "H" (时), 默认 "no"， 不切分
		JsonFormat: false,                                                                       // 写入文件的数据是否 json 格式化
		Format:     "[%timestamp_format%] [%level_string%] [%file% %function% %line%] [%body%]", // 如果写入文件的数据不 json 格式化，自定义日志格式
	}
	// 添加 file 为 Logger 的一个输出
	Logger.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig)

}
