
// 模块
/*
	脚手架
		mqant以模块化来组织代码模块，模块概念在框架中非常重要

	模块定义
		结构体只要实现了以上几个函数就被认为是一个模块

		指定一个模块的名称，非常重要，在配置文件和RPC路由中会频繁使用
		func GetType() string

		指定模块的版本
		func Version() string

		模块的初始化函数，当框架初始化模块是会主动调用该方法
		func OnInit(app module.App, settings *conf.ModuleSettings)

		当App解析配置后调用，这个接口不管这个模块是否在这个进程的模块分组中都会调用
		func OnAppConfigurationLoaded(app module.App)

		模块独立运行函数，框架初始化模块以后会以单独goroutine运行该函数，并且提供一
		个关闭信号，以在框架要停止模块时通知
		func Run(closeSig chan bool)

		当模块停止运行后框架会调用该方法，让模块做一些回收操作
		func OnDestory()

	模块生命周期

					mqant start														mqant stop	
						|																|
						|																|
			根据ProcessID过滤本进												  <————发送停止信号
			  程需要运行的module												  |		   |
						|													   |		|		
						|												    closeSig    |
					OnInit()												   |      finish
						|													   |		|
						|													   |		|
					  Run()		———goroutine———> Run(closeSig chan bool)  <————		onDestory()
																						|
																						|
																					  mqant end
																					
	模块使用
		通常我们不止实现一个简单模块，还需要利用框架的其他高级特性，因此我们通常会继承框架封装好的一些基础模块

		RPCModule
			继承basemodule.BaseModule该模块封装了mqant的RPC通信相关方法

		GateModule
			继承basegate.Gate该模块封装了tcp/websocket+mqtt协议的长连接网关

	不在进程分组中的模块如何初始化
		func (self *HelloWorld) OnAppConfigurationLoaded(app module.App) {
			// 当App初始化调用，这个接口不管这个模块是否在这个进程运行都会调用
			self.BaseModule.OnAppConfigurationLoaded(app)
		}
*/