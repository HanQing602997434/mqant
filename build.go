
// 搭建
/*
	安装mqant
	
	依赖
		go 1.13

		require github.com/liangdas/mqant v1.4.5

	代码组织结构
		工程目录
			|-bin
				|-conf
					|-server.json
			|-main.go

	入口
		main.go

		import (
			"github.com/liangdas/mqant"
			"github.com/liangdas/mqant/module"
		)

		func main() {
			app := mqant.CreateApp(
				module.Debug(true), // 只有是在调式模式下才会在控制台打印日志，非调式模式下只在日志文件中输出日志
			)

			app.Run(
				// 已实现的模块都应该在此处传入
			)
		}

	运行
		如果此时直接运行会报错（没有指定配置文件）
		panic: config path error open xxx/bin/conf/server.json: no such file or directory

	配置
		先按目录创建一个配置文件模板
		server.json

		{
			"Module":{
			},
			"Mqtt":{
				"WriteLoopChanNum": 10,
				"ReadPackLoop": 1,
				"ReadTimeOut": 600,
				"WriteTimeOut": 300
			},
			"Rpc":{
				"MaxCoroutine": 10000,
				"RpcExpired": 1,
				"LogSuccess": false
			}
		}

	运行
		2020-05-03T18:00:55.033228+08:00 [-] [-] [] [E] nats agent: nats: no servers available for connection
		Server configuration path : //work/go/mqant-helloworld/bin/conf/server.json
		2020-05-03T18:00:55.034615+08:00 [-] [-] [development] [I] [app.go:203] mqant  starting up
		2020-05-03T18:00:55.034707+08:00 [-] [-] [development] [I] [ModuleManager.go:50] This service ModuleGroup(ProcessID) is [development]

		可以看到框架已经运行起来了，但是什么都没有，仅仅运行了一个脚手架而已
*/