package main

import (
	"comment/core"
	"comment/flag"
	"comment/global"
	"comment/routers"
	"fmt"
	"google.golang.org/grpc"
)

// @title server API文档
// @version 1.0
// @description server API文档
// @host 127.0.0.1:9190
// @BasePath /
func main() {
	// 读取配置文件
	core.UMYaml()
	// 初始化日志
	global.Log = core.InitLogger()
	// 初始化数据库
	core.GrpcServer()
	global.DB = core.Gorm()
	global.Redis = core.Redis()
	global.GrpcConn = core.GrpcClient()
	core.Kafka()

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			global.Log.Error(err)
		}
	}(global.GrpcConn)

	//// 初始化es
	//global.Elasticsearch = core.ESInit()

	// 命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	r := routers.InitRouters()
	serAddr := global.Config.System.Addr()
	global.Log.Info("server run addr:", serAddr)
	err := r.Run(serAddr)
	if err != nil {
		return
	}
}
