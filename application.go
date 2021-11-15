
// 应用
/*
	模块
		应用(app)是mqant的最基本单位，通常一个进程中只需要实例化一个应用(app)。应用负责维护整个框架的
		基本服务
			1.服务注册与发现

			2.RPC通信

			3.模块依赖

	生命周期
		
					mqant Run
						|						——————>  mqant stop	
						|						|			 |
					初始化配置					|			 |
						|						|		  停止模块
						|						|			 |
			app.configurationLoaded()			|			 |
						|						|	等待指定时间后如模块还
						|						|	未停止成功则强杀进程
		 mods.OnAppConfigurationLoaded()		|			 |
						|						|			 |
						|						|		 mqant end
					初始化模块					|
						|						|
						|						|
				  app.startup()			 ———————>

	配置解析完成
		_ = app.OnConfigurationLoaded(func(app module.App) {

		})

	应用启动完成
		包括模块启动完成
		app.OnStartup(func (app module.App) {

		})

	设置强杀时间
		当出现模块卡死等无法退出进情况下，超过设置时间会强杀
		app := mqant.CreateApp(
			module.KillWaitTTL(1 * time.Minute)
		)
*/	