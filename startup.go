
// 启动参数
/*
	mqant默认解析参数
		mqant默认会解析启动环境变量，即调用flag.Parse()，如不想mqant解析可启动方法module.Parse(false)关闭

	mqant解析字段
		wdPath = *flag.String("wd", "", "Server work directory")
		confPath = *flag.String("conf", "", "Server configuration file path")
		ProcessID = *flag.String("pid", "development", "Server ProcessID")
		Logdir = *flag.String("log", "", "Log file directory")
		BIdir = *flag.String("bi", "", "bi file directory?")

	关闭mqant解析
		app := mqant.CreateApp(
			module.Parse(false), // 关闭后mqant所需参数需设置
		)

	指定进程工作路径

	启动命令设置
		
		module.Parse(true)
		命令wd

		mqant-example -wd /my/workdir

	初始化设置
		
		module.Parse(false)

		app := mqant.CreateApp()

	工作路径
		mqant会在工作路径上初始化未指定的设置

		配置文件{workdir}/bin/conf/server.json
		日志文件目录{workdir}/bin/logs
		BI日志文件目录{workdir}/bin/bi

	指定配置文件

	启动命令设置
		
		module.Parse(true)
		命令conf

		mqant-example -conf /my/config.json

	初始化设置
		
		module.Parse(false)
			app := mqant.CreateApp(
				module.Parse(false),
				module.Configure("/my/config.json")
			)

	指定模块分组ID

	启动命令设置
	
		module.Parse(true)
		命令 pid

		mqant-example -pid myPid

	初始化设置

		module.Parse(false)
			app := mqant.CreateApp(
				module.Parse(false),
				module.ProcessID("myPid"),
			)
*/