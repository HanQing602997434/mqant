
// mqant框架介绍
/*
	mqant是一个微服务框架。目标是简化分布式系统开发。mqant的核心是简单易用、关注业务场景，因此会针对
	特定场景研究一些特定组件和解决方案，方便开发者使用。

	mqant核心组件组成
		核心RPC组件 — 它提供了用于服务发现，客户端负载平衡，编码，同步通信库。

		http网关 — 提供将HTTP请求路由到相应微服务的API网关。它充当单个入口点，可以用作反向代理或将
		HTTP请求转换为RPC。

		tcp/websocket网关 — 它提供了tcp和websocket两种客户端连接方式。并且默认内置了一个简单的mqtt协
		议，可以非常便捷的提供长连接服务，快速搭建iot(物联网)和网络游戏通信业务，同时也支持自定义通信协
		议插件。

	名称定义
		模块 — 功能的统称，在mqant中以模块划分代码结构。
		
		服务 — 服务即模块，在微服务中服务比模块更容易让人理解。

		节点 — 即模块（服务）部署后的实例，一个模块（服务）可以在多台服务器启动多个节点。
*/