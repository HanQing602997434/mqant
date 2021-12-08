
// 服务发现
/*
	概述
		服务发现配置只能通过启动代码设置，包含
			nats配置
			注册中心配置（consul，etcd）
			服务发现TTL和注册间隔

	示例
		app := mqant.CreateApp(
    		module.Debug(true),  //只有是在调试模式下才会在控制台打印日志, 非调试模式下只在日志文件中输出日志
   			module.Nats(nc),     //指定nats rpc
    		module.Registry(rs), //指定服务发现
    		module.RegisterTTL(20*time.Second),
    		module.RegisterInterval(10*time.Second),
		)
*/