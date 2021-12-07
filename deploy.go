
// 部署概述
/*
	概述
		mqant部署分为单机部署和分布式部署，通常情况下，项目的所有模块代码都被编译到一个可执行文件中。
		在分布式部署时，我们通常想将网关模块跟后端服务模块分服务器部署，即：

			网关服务器仅启用网关模块
			后端服务器仅启动后端模块

	模块分组(ProcessID)
		模块分组便是为了实现上面的功能而设计的，如果要不同的模块分开部署可以按如下步骤操作

			在配置文件中将模块的ProcessID分开
			在启动应用进程时指定进程ProcessID

	单机部署
		mqant默认的模块分组值约定为development
		在调式期间可以将所有模块的分组ID都设置为development，这样一个进程就可以启用所有已实现的模块

	模块ProcessID设置

		"Module":{
			"moduletype":[
				{
					"Id":"moduletype",
					"ProcessID":"development"
				}
			]
		}

	指定进程ProcessID
		pid := flag.String("pid", "", "Server work directory")
		flag.Parse() // 解析输入的参数
		app := mqant.CreateApp(
			module.Debug(true),
			module.Parse(false),
			module.ProcessID(*pid),
		)

	编译
		sh build.sh

	运行
		mqant-example -pid mypid
		
			Server configuration path : /work/go/mqant-example/bin/conf/server.json
			2020-05-05T17:15:46.48651+08:00 [-] [-] [mypid] [I] [app.go:209] mqant  starting up
			2020-05-05T17:15:46.486779+08:00 [-] [-] [mypid] [I] [ModuleManager.go:50] This service ModuleGroup(ProcessID) is [mypid]
*/