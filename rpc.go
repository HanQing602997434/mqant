
// RPC
/*
	mqant RPC本身是一个相对独立的功能，RPC有以下的几个特点：
		1.目前支持nats作为服务发现通道，理论上可以扩展其他通信方式
		2.支持服务注册发现，是一个相对完善的微服务框架

	在模块中使用RPC
		module.BaseModule中已经集成了RPC，使用方式如下

	服务提供者
		注册handler

		注册服务函数
		module.GetServer().RegisterGO(_func string, fn interface{})

		注册服务函数
		module.GetServer().Register(_func string, fn interface{})

		RegisterGO和Register的区别是前者为每一条消息创建一个单独的协程来处理，后者注册的函数共用一个
		协程来处理所有消息，具体使用哪一种方式可以根据实际情况来定，但Register方式的函数一定注意不要
		执行耗时功能，以免引起消息阻塞

	服务调用者
		在开发过程中，模块A可能需要用到模块B的服务，这时模块A就成为了服务调用方。mqant提供了多种RPC
		调用方法，也支持高级扩展（服务发现）
*/