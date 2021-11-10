
// 架构
/*
	现如今只有多进程的架构才能达到支持较多在线用户，降低服务器压力，降低单点故障所带来的影响等要求，
	因此一个真正高可扩展的游戏运行架构必须是多进程的。

	然而在游戏的开发和运营也是按步骤阶段性进行的，尤其是现如今服务器硬件设备配置也越来越高的前提下，
	在游戏刚开始运营时单台服务器就足够支撑了，况且多进程部署所带来的运维成本也相对较高。

	mqant的设计思想是在能用单台服务器时能充分挖掘服务器的性能，而在需要多进程时再通过简单的配置就
	可以实现分布式部署。

	mqant框架架构图

		业务组件          游戏房间模型   短信验证码   项目业务模块

		mqant核心组件	  tcp/ws网关   http网关   mqtt协议库   自定义协议机制   日志库   ......

		mqant核心RPC框架  服务发现   负载平衡，编码   nats同步通信库

	mqant模块化运行架构图

		HD_Register (frontend)   UserManager(module)   mqant   Battle(module)			HD_CreateRoom (backend)
		用户注册				  ServerId : UM:01				ServerId : Battle:01	 创建房间
								 protocol : RPC				   protocol : RPC			...
										  	 \						   /
										   	  \					   	  /
											   \	 				 /
											    \	 			    /
												RPC(rabbitmq/go chan)
											   /         |          \
											 /			 |            \
										   /			 |              \
										 /               |                \
							Gate1(module)		 Gate2(module)		    Gate2(module)
							ServerId : Gate:01   ServerId : Gate:02     ServerId : Gate:03
							protocol : MQTT		 protocol : MQTT		protocol : MQTT
							/	   \				/         \				   |				
						   /        \              /           \               |
						用户01	用户02		用户03         用户04		  用户05
						IOS(tcp)  Android(tcp) H5(websocket)  Android(tcp)  Android(tcp)

	处理器（handler）

		handler就是一个可供远程RPC调用的函数

		每一个模块可以注册多个处理器（handler），例如用户模块可以提供用户注册、用户登录、用户注销、用户信息
		查询等一系列的处理器

	模块间通信RPC
		
		mqant模块间的通信方式应该使用RPC，而非本地代码调用。mqant的RPC也非才去grpc等消息投递工具，而
		是选用了nats消息总线。

	为什么选择消息队列进行RPC通信？
		选择消息队列而不选择传统的tcp/socket，rpc的主要考虑时传统的基于点对service/client模式的连接比较难
		于维护和统计，假如服务器存在100个模块和服务器的话进一个进程所需要维护的clinet连接就>100。

		而选择消息队列的话每个进程对每个模块只需要维护一条连接即可，nats有完善的监控，报警工具，可以随时监控
		模块的处理性能和实时性。
*/